package dockerclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types/swarm"
	goclient "github.com/fsouza/go-dockerclient"
)

const (
	defaultHttpRequestTimeout = time.Second * 10
)

type DockerGoClient struct {
	Client   DockerClient
	Endpoint string
}

func (dg *DockerGoClient) Ping() error {
	return dg.Client.Ping()
}

func NewDockerGoClient(endpoint string) (*DockerGoClient, error) {
	client, err := goclient.NewVersionedClient(endpoint, "1.23")

	if err != nil {
		log.Error("Unable to connect to docker daemon . Ensure docker is running endpoint ", endpoint, "err", err)
		return nil, err
	}

	// Even if we have a dockerclient, the daemon might not be running. Ping it
	// to ensure it's up.
	//err = client.Ping()
	//if err != nil {
	//	log.Error("Unable to ping docker daemon. Ensure docker is running endpoint ", endpoint, "err", err)
	//	return nil, err
	//}

	return &DockerGoClient{
		Client:   client,
		Endpoint: endpoint,
	}, nil
}

// NodeList returns the list of nodes.
func (dg *DockerGoClient) Nodelist(opts NodeListOptions) ([]swarm.Node, error) {
	var nodes []swarm.Node
	targetUrl, err := url.Parse(dg.Endpoint + "/nodes")
	if err != nil {
		return nodes, err
	}

	log.Info(targetUrl.String())
	content, err := executeHttpGet(*targetUrl)
	if err != nil {
		return nodes, err
	}

	if err := json.Unmarshal(content, &nodes); err != nil {
		return nodes, err
	}

	return nodes, nil
}

// execute http get request use default timeout
func executeHttpGet(targetUrl url.URL) ([]byte, error) {
	client := http.Client{Timeout: defaultHttpRequestTimeout}
	resp, err := client.Get(targetUrl.String())
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http response status code is %d not 200", resp.StatusCode)
	}

	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
