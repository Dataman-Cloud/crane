package dockerclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Executes http GET request with default timeout.
func (client *RolexDockerClient) HttpGet(requestUrl string, query url.Values, headers map[string][]string) ([]byte, error) {
	apiPath := getAPIPath(requestUrl, query)
	resp, err := client.sharedHttpClient.Get(apiPath)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http response status code is %d error: %s", resp.StatusCode, string(result))
	}

	return result, nil
}

// Executes http DELETE request with default timeout.
func (client *RolexDockerClient) HttpDelete(requestUrl string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", requestUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.sharedHttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("http response status code is %d error: %s", resp.StatusCode, string(result))
	}

	return result, nil
}

// Executes http POST request with default timeout.
func (client *RolexDockerClient) HttpPost(requestUrl string, query url.Values, obj interface{}, headers map[string][]string) ([]byte, error) {
	apiPath := getAPIPath(requestUrl, query)

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

	req, err := http.NewRequest("POST", apiPath, body)

	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header[k] = v
		}
	}

	resp, err := client.sharedHttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("http response status code is %d error: %s", resp.StatusCode, string(result))
	}

	return result, nil
}

// Executes http PUT request with default timeout.
func (client *RolexDockerClient) HttpPut(requestUrl string, query url.Values, body []byte, headers map[string][]string) ([]byte, error) {
	apiPath := getAPIPath(requestUrl, query)
	req, err := http.NewRequest("PUT", apiPath, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if headers != nil {
		for k, v := range headers {
			req.Header[k] = v
		}
	}

	resp, err := client.sharedHttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http response status code is %d error: %s", resp.StatusCode, string(result))
	}

	return result, nil
}

// getAPIPath returns the versioned request path to call the api.
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
