package httpclient

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
)

const (
	DefaultTimeout = time.Second * 15
)

type Client struct {
	HttpClient        *http.Client
	TLSConfig         *tls.Config
	CustomHTTPHeaders map[string]string
}

//TODO learn more from https://github.com/hashicorp/go-cleanhttp
// DefaultTransport returns a new http.Transport with the same default values
// as http.DefaultTransport, but with idle connections and keepalives disabled.
func DefaultTransport() *http.Transport {
	transport := DefaultPooledTransport()
	transport.DisableKeepAlives = true
	transport.MaxIdleConnsPerHost = -1
	return transport
}

// DefaultPooledTransport returns a new http.Transport with similar default
// values to http.DefaultTransport. Do not use this for transient transports as
// it can leak file descriptors over time. Only use this for transports that
// will be re-used for the same host(s).
func DefaultPooledTransport() *http.Transport {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		DisableKeepAlives:   false,
		MaxIdleConnsPerHost: 1,
	}
	return transport
}

// DefaultClient returns a new http.Client with similar default values to
// http.Client, but with a non-shared Transport, idle connections disabled, and
// keepalives disabled.
func DefaultClient() *http.Client {
	return &http.Client{
		Transport: DefaultTransport(),
	}
}

// DefaultPooledClient returns a new http.Client with the same default values
// as http.Client, but with a shared Transport. Do not use this function
// for transient clients as it can leak file descriptors over time. Only use
// this for clients that will be re-used for the same host(s).
func DefaultPooledClient() *http.Client {
	return &http.Client{
		Transport: DefaultPooledTransport(),
	}
}

func NewClient(httpClient *http.Client, customHeaders map[string]string) (*Client, error) {
	if httpClient == nil {
		httpClient = DefaultClient()
		httpClient.Timeout = DefaultTimeout
	}

	return &Client{
		HttpClient:        httpClient,
		CustomHTTPHeaders: customHeaders,
	}, nil
}

func NewTLSClient(tlsCaCert, tlsCert, tlsKey string, httpClient *http.Client, customHeaders map[string]string) (*Client, error) {
	caCert, err := ioutil.ReadFile(tlsCaCert)
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	httpTLSCert, err := tls.LoadX509KeyPair(tlsCert, tlsKey)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{httpTLSCert},
		RootCAs:      caCertPool,
	}

	tlsConfig.BuildNameToCertificate()

	if httpClient == nil {
		httpClient = DefaultClient()
		httpClient.Timeout = DefaultTimeout
	}

	tr := DefaultTransport()
	tr.TLSClientConfig = tlsConfig
	httpClient.Transport = tr
	return &Client{
		HttpClient:        httpClient,
		TLSConfig:         tlsConfig,
		CustomHTTPHeaders: customHeaders,
	}, nil
}

func (cli *Client) POST(ctx context.Context, requestUrl string, query url.Values, obj interface{}, headers map[string][]string) ([]byte, error) {
	resp, err := cli.do(ctx, requestUrl, "POST", query, obj, headers)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("%s", parseMessage(result))
	}

	return result, nil
}

func (cli *Client) PUT(ctx context.Context, requestUrl string, query url.Values, obj interface{}, headers map[string][]string) ([]byte, error) {
	resp, err := cli.do(ctx, requestUrl, "PUT", query, obj, headers)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", parseMessage(result))
	}

	return result, nil
}

func (cli *Client) DELETE(ctx context.Context, requestUrl string, query url.Values, headers map[string][]string) ([]byte, error) {
	resp, err := cli.do(ctx, requestUrl, "DELETE", query, nil, headers)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("%s", parseMessage(result))
	}

	return result, nil
}

func (cli *Client) GET(ctx context.Context, requestUrl string, query url.Values, headers map[string][]string) ([]byte, error) {
	resp, err := cli.do(ctx, requestUrl, "GET", query, nil, headers)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", parseMessage(result))
	}

	return result, nil
}

func (cli *Client) PATCH(ctx context.Context, requestUrl string, query url.Values, obj interface{}, headers map[string][]string) ([]byte, error) {
	resp, err := cli.do(ctx, requestUrl, "PATCH", query, obj, headers)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", parseMessage(result))
	}

	return result, nil
}

func (cli *Client) do(ctx context.Context, requestUrl, method string, query url.Values, obj interface{}, headers map[string][]string) (*http.Response, error) {
	var body io.Reader

	if obj != nil {
		var err error
		body, err = encodeData(obj)
		if err != nil {
			return nil, err
		}
		if headers == nil {
			headers = make(map[string][]string)
		}
		headers["Content-Type"] = []string{"application/json"}
	}

	return cli.sendRequest(ctx, requestUrl, method, query, body, headers)
}

func (cli *Client) sendRequest(ctx context.Context, requestUrl, method string, query url.Values, body io.Reader, headers map[string][]string) (*http.Response, error) {
	expectedPayload := (method == "POST" || method == "PUT" || method == "PATCH")
	if expectedPayload && body == nil {
		body = bytes.NewReader([]byte{})
	}

	req, err := cli.newRequest(method, requestUrl, query, body, headers)
	if err != nil {
		return nil, err
	}

	if expectedPayload && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "text/plain")
	}

	if ctx == nil {
		ctx = context.Background()
	}

	resp, err := ctxhttp.Do(ctx, cli.HttpClient, req)
	//TODO classify error return more info
	if err != nil {
		return nil, chooseError(ctx, err)
	}

	return resp, nil
}

// if error in context, return that instead of generic http error
func chooseError(ctx context.Context, err error) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return err
	}
}

func (cli *Client) newRequest(method, path string, query url.Values, body io.Reader, headers map[string][]string) (*http.Request, error) {
	apiPath := getAPIPath(path, query)
	req, err := http.NewRequest(method, apiPath, body)
	if err != nil {
		return nil, err
	}

	for k, v := range cli.CustomHTTPHeaders {
		req.Header.Set(k, v)
	}

	if headers != nil {
		for k, v := range headers {
			req.Header[k] = v
		}
	}

	return req, nil
}

// returns the versioned request path to call the api.
// It appends the query parameters to the path if they are not empty.
func getAPIPath(apiPath string, query url.Values) string {
	u := &url.URL{
		Path: apiPath,
	}
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}
	return u.String()
}

func encodeData(data interface{}) (*bytes.Buffer, error) {
	params := bytes.NewBuffer(nil)
	if data != nil {
		if err := json.NewEncoder(params).Encode(data); err != nil {
			return nil, err
		}
	}
	return params, nil
}

func parseMessage(result []byte) string {
	var r struct {
		Message string `json:"message"`
	}
	err := json.Unmarshal(result, &r)
	if err != nil {
		return string(result)
	}

	return r.Message
}
