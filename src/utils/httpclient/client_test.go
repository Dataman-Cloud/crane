package httpclient

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestDefaultTransport(t *testing.T) {
	assert.NotNil(t, DefaultTransport())
}

func TestDefaultPooledTransport(t *testing.T) {
	assert.NotNil(t, DefaultPooledTransport())
}

func TestDefaultClient(t *testing.T) {
	assert.NotNil(t, DefaultClient())
}

func TestNewClient(t *testing.T) {
	client := http.Client{}
	nClient, err := NewClient(&client, make(map[string]string))
	assert.NotNil(t, nClient)
	assert.Nil(t, err)
}

func NewTestClient() *Client {
	httpClient := http.Client{}
	c, _ := NewClient(&httpClient, make(map[string]string))
	return c
}

type Message struct {
	Message string `json:"message"`
}

var msg Message = Message{Message: "foobar"}

func NewTestServer() *httptest.Server {
	body, _ := json.Marshal(msg)

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, string(body))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))

	return server
}

func TestGet(t *testing.T) {
	server := NewTestServer()
	defer server.Close()
	c := context.Background()

	result, err := NewTestClient().GET(c, server.URL, url.Values{}, make(map[string][]string))
	assert.Nil(t, err)
	b, _ := json.Marshal(msg)
	assert.Equal(t, b, result)
}

func TestPOST(t *testing.T) {
	server := NewTestServer()
	defer server.Close()
	c := context.Background()

	result, err := NewTestClient().POST(c, server.URL, url.Values{}, nil, make(map[string][]string))
	assert.Nil(t, err)
	b, _ := json.Marshal(msg)
	assert.Equal(t, b, result)
}

func TestPUT(t *testing.T) {
	server := NewTestServer()
	defer server.Close()
	c := context.Background()

	result, err := NewTestClient().PUT(c, server.URL, url.Values{}, nil, make(map[string][]string))
	assert.Nil(t, err)
	b, _ := json.Marshal(msg)
	assert.Equal(t, b, result)
}

func TestDELETE(t *testing.T) {
	server := NewTestServer()
	defer server.Close()
	c := context.Background()

	result, err := NewTestClient().DELETE(c, server.URL, url.Values{}, make(map[string][]string))
	assert.Nil(t, err)
	b, _ := json.Marshal(msg)
	assert.Equal(t, b, result)
}

func TestPATCH(t *testing.T) {
	server := NewTestServer()
	defer server.Close()
	c := context.Background()

	result, err := NewTestClient().PATCH(c, server.URL, url.Values{}, nil, make(map[string][]string))
	assert.Nil(t, err)
	b, _ := json.Marshal(msg)
	assert.Equal(t, b, result)
}

func TestParseMessage(t *testing.T) {
	b, _ := json.Marshal(msg)
	assert.Equal(t, "foobar", parseMessage(b))
}

func TestEncodeBuffer(t *testing.T) {
	b, err := encodeData(msg)
	assert.Nil(t, err)
	assert.NotNil(t, b)
}

func TestGetApiPath(t *testing.T) {
	urlValues := url.Values{}

	urlValues.Set("len", "1")
	assert.Equal(t, "http://localhost?len=1", getAPIPath("http://localhost", urlValues))

	assert.Equal(t, "http://localhost", getAPIPath("http://localhost", nil))
}

func TestChooseError(t *testing.T) {
	c := context.Background()
	var err error
	assert.Equal(t, err, chooseError(c, err))

	c = context.Background()
	ctx, cancel := context.WithCancel(c)
	cancel()

	assert.NotNil(t, chooseError(ctx, nil))
}

func TestSendRequest(t *testing.T) {
	server := NewTestServer()
	defer server.Close()
	c := context.Background()

	resp, err := NewTestClient().sendRequest(c, server.URL, "GET", url.Values{}, nil, make(map[string][]string))
	assert.Nil(t, err)
	defer resp.Body.Close()
	assert.NotNil(t, resp)
}

func TestDo(t *testing.T) {
	server := NewTestServer()
	defer server.Close()
	c := context.Background()

	resp, err := NewTestClient().do(c, server.URL, "GET", url.Values{}, nil, make(map[string][]string))
	assert.Nil(t, err)
	defer resp.Body.Close()
	assert.NotNil(t, resp)
}
