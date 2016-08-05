package dockerclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Executes an http GET request with default timeout.
func (client *RolexDockerClient) HttpGet(requestUrl string, query url.Values, headers map[string][]string) ([]byte, error) {
	apiPath := client.getAPIPath(requestUrl, query)
	resp, err := client.SharedHttpClient.Get(apiPath)
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

// execute http delete request use default timeout
func (client *RolexDockerClient) HttpDelete(requestUrl string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", requestUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.SharedHttpClient.Do(req)
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

func (client *RolexDockerClient) HttpPost(requestUrl string, query url.Values, body []byte, headers map[string][]string) ([]byte, error) {
	apiPath := client.getAPIPath(requestUrl, query)
	req, err := http.NewRequest("POST", apiPath, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if headers != nil {
		for k, v := range headers {
			req.Header[k] = v
		}
	}

	resp, err := client.SharedHttpClient.Do(req)
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

func (client *RolexDockerClient) HttpPut(requestUrl string, query url.Values, body []byte, headers map[string][]string) ([]byte, error) {
	apiPath := client.getAPIPath(requestUrl, query)
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

	resp, err := client.SharedHttpClient.Do(req)
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

// It looks like this function doesn't need a receiver.
// getAPIPath returns the versioned request path to call the api.
// It appends the query parameters to the path if they are not empty.
func (client *RolexDockerClient) getAPIPath(apiPath string, query url.Values) string {
	u := &url.URL{
		Path: apiPath,
	}
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}
	return u.String()
}
