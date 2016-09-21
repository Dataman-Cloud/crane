package tty

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"text/template"
	"unsafe"

	log "github.com/Sirupsen/logrus"
	"github.com/fatih/structs"
	"github.com/gorilla/websocket"
	"github.com/kr/pty"
)

const (
	SubProtocol = "shurenyun"
)

type ClientContext struct {
	Request       *http.Request
	Conn          *websocket.Conn
	Command       *exec.Cmd
	Pty           *os.File
	WriteMutex    *sync.Mutex
	TitleTemplate *template.Template
	Options       *Options
}

type Options struct {
	CloseSignal       int                    `hcl:"close_signal"`
	Preferences       HtermPrefernces        `hcl:"preferences"`
	RawPreferences    map[string]interface{} `hcl:"preferences"`
	EnableReconnect   bool                   `hcl:"enable_reconnect"`
	ReconnectTime     int                    `hcl:"reconnect_time"`
	Once              bool                   `hcl:"once"`
	PermitWrite       bool                   `hcl:"permit_write"`
	TitleFormat       string                 `hcl:"title_format"`
	EnableCustomTitle bool                   `hcl:"custom_title"`
	EnableRandomUrl   bool                   `hcl:"enable_random_url"`
	RandomUrlLength   int                    `hcl:"random_url_length"`
}

const (
	Input          = '0'
	Ping           = '1'
	ResizeTerminal = '2'
)

const (
	Output         = '0'
	Pong           = '1'
	SetWindowTitle = '2'
	SetPreferences = '3'
	SetReconnect   = '4'
)

type ArgResizeTerminal struct {
	Columns float64
	Rows    float64
}

type ContextVars struct {
	Command    string
	Pid        int
	Hostname   string
	RemoteAddr string
}

var (
	Upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		Subprotocols:    []string{SubProtocol},
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	DefaultOptions = &Options{
		PermitWrite:       true,
		EnableRandomUrl:   false,
		RandomUrlLength:   8,
		TitleFormat:       SubProtocol + "TTY - {{ .Command }} ({{ .Hostname }})",
		EnableCustomTitle: false,
		EnableReconnect:   false,
		ReconnectTime:     10,
		Once:              false,
		CloseSignal:       1, // syscall.SIGHUP
		Preferences:       HtermPrefernces{},
	}
)

func New(cmd *exec.Cmd, conn *websocket.Conn, request *http.Request, opts *Options) (*ClientContext, error) {
	ptyIo, err := pty.Start(cmd)
	if err != nil {
		return nil, err
	}

	titleTemplate, err := template.New("title").Parse(opts.TitleFormat)
	if err != nil {
		return nil, errors.New("Title format string syntax error")
	}

	return &ClientContext{
		Request:       request,
		Conn:          conn,
		Command:       cmd,
		Pty:           ptyIo,
		TitleTemplate: titleTemplate,
		WriteMutex:    &sync.Mutex{},
		Options:       opts,
	}, nil
}

func (ctx *ClientContext) HandleClient() {
	log.Infof("Command is running fot client %s with PID %d (args=%q)",
		ctx.Request.RemoteAddr, ctx.Command.Process.Pid, ctx.Command.Args)
	exit := make(chan bool, 2)

	go func() {
		defer func() { exit <- true }()

		ctx.processSend()
	}()

	go func() {
		defer func() { exit <- true }()

		ctx.processReceive()
	}()

	go func() {
		<-exit

		// Even if the Pty has been closed
		// Read(0 processSend() keeps blocking and the process doen't exit
		ctx.Command.Process.Signal(syscall.Signal(ctx.Options.CloseSignal))
		ctx.Command.Wait()

		ctx.Pty.Close()

		ctx.Conn.Close()
		log.Infof("Connection closed: %s", ctx.Request.RemoteAddr)
	}()
}

func (ctx *ClientContext) processSend() {
	if err := ctx.sendInitialize(); err != nil {
		log.Error("Init websocket conn got error: ", err)
		return
	}

	buf := make([]byte, 1024)

	for {
		size, err := ctx.Pty.Read(buf)
		if err != nil {
			log.Errorf("Command exited of %s. Error: %s", ctx.Request.RemoteAddr, err.Error())
			return
		}

		safeMessage := base64.StdEncoding.EncodeToString([]byte(buf[:size]))
		if err := ctx.write(append([]byte{Output}, []byte(safeMessage)...)); err != nil {
			log.Error("Websocket send message got error: ", err)
			return
		}
	}
}

func (ctx *ClientContext) write(data []byte) error {
	ctx.WriteMutex.Lock()
	defer ctx.WriteMutex.Unlock()
	return ctx.Conn.WriteMessage(websocket.TextMessage, data)
}

func (ctx ClientContext) sendCustomTitle() error {
	hostname, _ := os.Hostname()
	titleVars := ContextVars{
		Command:    strings.Join(ctx.Command.Args, " "),
		Pid:        ctx.Command.Process.Pid,
		Hostname:   hostname,
		RemoteAddr: ctx.Request.RemoteAddr,
	}

	titleBuffer := new(bytes.Buffer)
	if err := ctx.TitleTemplate.Execute(titleBuffer, titleVars); err != nil {
		log.Error("Create app title got error: ", err)
		return err
	}

	if err := ctx.write(append([]byte{SetWindowTitle}, titleBuffer.Bytes()...)); err != nil {
		log.Error("Write window title got error: ", err)
		return err
	}

	return nil
}

func (ctx ClientContext) sendInitialize() error {
	if ctx.Options.EnableCustomTitle {
		if err := ctx.sendCustomTitle(); err != nil {
			return err
		}
	}
	prefStruct := structs.New(ctx.Options.Preferences)
	prefMap := prefStruct.Map()
	htermPrefs := make(map[string]interface{})

	for key, value := range prefMap {
		rawKey := prefStruct.Field(key).Tag("hcl")
		if _, ok := ctx.Options.RawPreferences[rawKey]; ok {
			htermPrefs[strings.Replace(rawKey, "_", "-", -1)] = value
		}
	}

	prefs, err := json.Marshal(htermPrefs)
	if err != nil {
		log.Error("Marshal htermPrefs got error: ", err)
		return err
	}

	if err := ctx.write(append([]byte{SetPreferences}, prefs...)); err != nil {
		log.Error("Write prefs got error: ", err)
		return err
	}

	if ctx.Options.EnableReconnect {
		reconnect, _ := json.Marshal(ctx.Options.ReconnectTime)
		if err := ctx.write(append([]byte{SetReconnect}, reconnect...)); err != nil {
			log.Error("Write reconnect info got error: ", err)
			return err
		}
	}

	return nil
}

func (ctx *ClientContext) processReceive() {
	for {
		_, data, err := ctx.Conn.ReadMessage()
		if err != nil {
			log.Error("Read message from websocket got error: ", err)
			return
		}

		if len(data) == 0 {
			log.Error("An error has occurred")
			return
		}

		switch data[0] {
		case Input:
			if !ctx.Options.PermitWrite {
				break
			}

			_, err := ctx.Pty.Write(data[1:])
			if err != nil {
				log.Error("Write received message got error: ", err)
				return
			}

		case Ping:
			if err := ctx.write([]byte{Pong}); err != nil {
				log.Error("Write message pong got error: ", err)
				return
			}

		case ResizeTerminal:
			var args ArgResizeTerminal
			if err := json.Unmarshal(data[1:], &args); err != nil {
				log.Error("Malformed remote command")
				return
			}

			window := struct {
				row uint16
				col uint16
				x   uint16
				y   uint16
			}{
				uint16(args.Rows),
				uint16(args.Columns),
				0,
				0,
			}

			syscall.Syscall(
				syscall.SYS_IOCTL,
				ctx.Pty.Fd(),
				syscall.TIOCSWINSZ,
				uintptr(unsafe.Pointer(&window)),
			)

		default:
			log.Error("Unknown message type")
			return
		}
	}
}
