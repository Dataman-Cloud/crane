package tty

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cmd := exec.Command("echo", "test")
	client, err := New(cmd, nil, nil, DefaultOptions)
	assert.Nil(t, err)

	assert.True(t, client.Options.PermitWrite)
	assert.False(t, client.Options.EnableRandomUrl)
	assert.Equal(t, client.Options.RandomUrlLength, 8)
	assert.Equal(t, client.Options.TitleFormat, SubProtocol+"TTY - {{ .Command }} ({{ .Hostname }})")
	assert.False(t, client.Options.EnableCustomTitle)
	assert.False(t, client.Options.EnableReconnect)
	assert.Equal(t, client.Options.ReconnectTime, 10)
	assert.False(t, client.Options.Once)
	assert.Equal(t, client.Options.CloseSignal, 1)
}

func startMockAcsServer(t *testing.T, closeWS <-chan bool) (*httptest.Server, <-chan error, error) {
	errChan := make(chan error, 1)

	upgrader := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		go func() {
			<-closeWS
			ws.WriteMessage(websocket.CloseMessage, nil)
			ws.Close()
			errChan <- io.EOF
		}()
		if err != nil {
			errChan <- err
		}

		req := &http.Request{
			RemoteAddr: "test.test",
		}
		cmd := exec.Command("ping", "www.baidu.com")
		client, err := New(cmd, ws, req, DefaultOptions)
		assert.Nil(t, err)
		client.HandleClient()
	})

	//server := httptest.NewTLSServer(handler)
	server := httptest.NewServer(handler)
	return server, errChan, nil
}

func parseURL(s string) string {
	var wssString string
	switch {
	case strings.HasPrefix(s, "http://"):
		wssString = strings.Replace(s, "http", "ws", -1)
	case strings.HasPrefix(s, "https://"):
		wssString = strings.Replace(s, "https", "wss", -1)
	default:
	}
	return wssString
}

//TODO: upccup make it right as soon
//func TestHandleClient(t *testing.T) {
//	closeWS := make(chan bool)
//	interrupt := make(chan os.Signal, 1)
//	signal.Notify(interrupt, os.Interrupt)
//	server, errs, err := startMockAcsServer(t, closeWS)
//	assert.Nil(t, err)
//
//	addr := parseURL(server.URL)
//	t.Log("server url: ", addr)
//
//	u, err := url.Parse(addr)
//	assert.Nil(t, err)
//
//	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
//	assert.Nil(t, err)
//	defer c.Close()
//
//	done := make(chan struct{})
//
//	go func() {
//		defer c.Close()
//		defer close(done)
//		for {
//			_, message, err := c.ReadMessage()
//			if err != nil {
//				log.Println("read:", err)
//				return
//			}
//			log.Printf("recv: %s", message)
//		}
//	}()
//
//	ticker := time.NewTicker(time.Second)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case t := <-ticker.C:
//			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
//			if err != nil {
//				log.Println("write:", err)
//				return
//			}
//		case <-interrupt:
//			log.Println("interrupt")
//			// To cleanly close a connection, a client should send a close
//			// frame and wait for the server to close the connection.
//			closeWS <- true
//			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
//			if err != nil {
//				log.Println("write close:", err)
//				return
//			}
//			select {
//			case <-done:
//			case <-time.After(time.Second):
//			}
//			c.Close()
//			return
//		case err := <-errs:
//			t.Error(err)
//			return
//		}
//	}
//}
