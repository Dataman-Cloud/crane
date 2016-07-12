package dockerclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// execute http get request use default timeout
func (client *RolexDockerClient) HttpGet(requestPath string) ([]byte, error) {
	resp, err := client.HttpClient.Get(client.HttpEndpoint + "/" + requestPath)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http response status code is %d", resp.StatusCode)
	}

	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// execute http delete request use default timeout
func (client *RolexDockerClient) HttpDelete(requestPath string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", client.HttpEndpoint+"/"+requestPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("http response status code is %d", resp.StatusCode)
	}

	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (client *RolexDockerClient) HttpPost(requestPath string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", client.HttpEndpoint+"/"+requestPath, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("http response status code is %d", resp.StatusCode)
	}

	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
