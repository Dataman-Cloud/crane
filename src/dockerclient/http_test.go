package dockerclient

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHttpGet(T *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test success"))
	}))
	defer server.Close()
	rolexDockerClient := &RolexDockerClient{
		sharedHttpClient: &http.Client{},
	}
	if _, err := rolexDockerClient.HttpGet(server.URL, url.Values{}, nil); err != nil {
		T.Error("httpget error")
	}
}

func TestHttpDelete(T *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test success"))
	}))
	defer server.Close()
	rolexDockerClient := &RolexDockerClient{
		sharedHttpClient: &http.Client{},
	}
	if _, err := rolexDockerClient.HttpDelete(server.URL); err != nil {
		T.Error("httpdelete error")
	}
}

func TestHttpPost(T *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test success"))
	}))
	defer server.Close()
	rolexDockerClient := &RolexDockerClient{
		sharedHttpClient: &http.Client{},
	}
	if _, err := rolexDockerClient.HttpPost(server.URL, url.Values{}, nil, nil); err != nil {
		T.Error("httppost error")
	}
}

func TestHttpPut(T *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test success"))
	}))
	defer server.Close()
	rolexDockerClient := &RolexDockerClient{
		sharedHttpClient: &http.Client{},
	}
	if _, err := rolexDockerClient.HttpPut(server.URL, url.Values{}, nil, nil); err != nil {
		T.Error("httpput error")
	}
}
func TestgetAPIPath(T *testing.T) {
	values := url.Values{}
	values.Add("test", "value")
	u := &url.URL{
		Path: "localhost",
	}
	if len(values) > 0 {
		u.RawQuery = values.Encode()
	}
	if tu := getAPIPath("localhost", values); u.String() != tu {
		T.Error("get api path error")
	}
}
