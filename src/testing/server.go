package mock

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"

	"github.com/gorilla/mux"
)

type Server struct {
	Addr    string
	Port    string
	Scheme  string
	server  *httptest.Server
	mux     *mux.Router
	routers []*RouterMap
}

type RouterMap struct {
	Path    string
	Method  string
	RGroups []*RGroup
}

type RGroup struct {
	Request  *Request
	Response *Response
}

type Request struct {
	Query      string
	BodyBuffer []byte
	Error      error
}

type Response struct {
	StatusCode int
	BodyBuffer []byte
	Error      error
}

func NewServer() *Server {
	mux := mux.NewRouter()
	server := httptest.NewServer(mux)
	parsedUrl, _ := url.Parse(server.URL)
	host, port, _ := net.SplitHostPort(parsedUrl.Host)
	return &Server{
		mux:    mux,
		Addr:   host,
		Port:   port,
		Scheme: parsedUrl.Scheme,
		server: server,
	}
}

func (s *Server) Close() {
	s.server.Close()
}

func (s *Server) Register() {
	for _, router := range s.routers {
		s.mux.Path(router.Path).Methods(router.Method).HandlerFunc(router.Handler)
	}
}

func (s *Server) AddRouter(path string, method string) *RouterMap {
	method = strings.ToUpper(method)
	rgroups := make([]*RGroup, 0)
	routerMap := &RouterMap{
		Path:    path,
		Method:  method,
		RGroups: rgroups,
	}
	s.routers = append(s.routers, routerMap)
	return routerMap
}

func (rm *RouterMap) RGroup() *RGroup {
	rGroup := &RGroup{
		Request:  &Request{},
		Response: &Response{},
	}
	rm.RGroups = append(rm.RGroups, rGroup)
	return rGroup
}

func (rg *RGroup) RBody(body io.Reader) *RGroup {
	rg.Request.BodyBuffer, rg.Request.Error = ioutil.ReadAll(body)
	return rg
}

func (rg *RGroup) RQuery(query string) *RGroup {
	rg.Request.Query = query
	return rg
}

func (rg *RGroup) RBodyString(body string) *RGroup {
	rg.Request.BodyBuffer = []byte(body)
	return rg
}

func (rg *RGroup) RFile(path string) *RGroup {
	rg.Request.BodyBuffer, rg.Request.Error = ioutil.ReadFile(path)
	return rg
}

func (rg *RGroup) RJSON(data interface{}) *RGroup {
	rg.Request.BodyBuffer, rg.Request.Error = readAndDecode(data, "json")
	return rg
}

func (rg *RGroup) Reply(status int) *RGroup {
	rg.Response.StatusCode = status
	return rg
}

func (rg *RGroup) WBody(body io.Reader) *RGroup {
	rg.Response.BodyBuffer, rg.Response.Error = ioutil.ReadAll(body)
	return rg
}

func (rg *RGroup) WBodyString(body string) *RGroup {
	rg.Response.BodyBuffer = []byte(body)
	return rg
}

func (rg *RGroup) WFile(path string) *RGroup {
	rg.Response.BodyBuffer, rg.Response.Error = ioutil.ReadFile(path)
	return rg
}

func (rg *RGroup) WJSON(data interface{}) *RGroup {
	rg.Response.BodyBuffer, rg.Response.Error = readAndDecode(data, "json")
	return rg
}

func readAndDecode(data interface{}, kind string) ([]byte, error) {
	buf := &bytes.Buffer{}

	switch data.(type) {
	case string:
		buf.WriteString(data.(string))
	case []byte:
		buf.Write(data.([]byte))
	default:
		var err error
		if kind == "xml" {
			err = xml.NewEncoder(buf).Encode(data)
		} else {
			err = json.NewEncoder(buf).Encode(data)
		}
		if err != nil {
			return nil, err
		}
	}

	return ioutil.ReadAll(buf)
}

func checkRawQuery(r *http.Request, want string) error {
	rQuery := r.URL.Query()
	wantQuery, err := url.ParseQuery(want)
	if err != nil {
		return errors.New(fmt.Sprintf("parse registered query fails: %v", want))
	}
	if reflect.DeepEqual(rQuery, wantQuery) == false {
		return errors.New(fmt.Sprintf("Request query is not equal : %v, want %v", rQuery, wantQuery))
	}
	return nil

}

func checkBody(r *http.Request, want []byte) error {
	rBody, _ := ioutil.ReadAll(r.Body)
	if bytes.Equal(rBody, want) == false {
		return errors.New("Request body is not equal")
	}
	return nil
}

func (rm *RouterMap) Handler(w http.ResponseWriter, r *http.Request) {
	isMatched := false
	for _, rGroup := range rm.RGroups {
		isMatched = false
		//checkQuery
		err1 := checkRawQuery(r, rGroup.Request.Query)
		if err1 != nil {
			continue
		}
		//check body
		err2 := checkBody(r, rGroup.Request.BodyBuffer)
		if err2 != nil {
			continue
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(rGroup.Response.StatusCode)
		w.Write(rGroup.Response.BodyBuffer)
		isMatched = true
		break
	}
	if !isMatched {
		http.Error(w, fmt.Sprintf(`{"message": there is no matched handler for request:%s}`, r.URL), 400)
	}
	return

}
