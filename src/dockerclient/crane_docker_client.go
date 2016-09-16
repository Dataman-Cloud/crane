package dockerclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/Dataman-Cloud/crane/src/utils/config"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/httpclient"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
)

type CraneDockerClient struct {
	DockerClientInterface

	// client connect to swarm cluster manager
	// TODO swarm cluster has multiple manager and can be changed at runtime
	// make sure swarmManager could connect to next if first one failed
	swarmManager *docker.Client

	// http client shared both for cluster connection & client connection
	sharedHttpClient         *httpclient.Client
	swarmManagerHttpEndpoint string

	config *config.Config
}

// initialize crane docker client
func NewCraneDockerClient(config *config.Config) (*CraneDockerClient, error) {
	var err error

	swarmManagerEntry := config.DockerEntryScheme + "://" + config.SwarmManagerIP + ":" + config.DockerEntryPort
	client := &CraneDockerClient{
		config: config,

		swarmManagerHttpEndpoint: swarmManagerEntry,
	}

	if config.DockerTlsVerify {
		client.swarmManager, err = NewGoDockerClientTls(swarmManagerEntry, API_VERSION, config)
		client.sharedHttpClient, err = NewHttpClientTls(config)
	} else {
		client.swarmManager, err = docker.NewVersionedClient(swarmManagerEntry, API_VERSION)
		client.sharedHttpClient, err = NewHttpClient()
	}

	if err != nil {
		log.Error("Unable to connect to docker daemon . Ensure docker is running endpoint ", swarmManagerEntry, "err: ", err)
		return nil, err
	}

	err = client.swarmManager.Ping()
	if err != nil {
		log.Error("Unable to ping docker daemon. Ensure docker is running endpoint ", swarmManagerEntry, "err: ", err)
		return nil, err
	}

	return client, nil
}

// return swarm docker client
func (client *CraneDockerClient) SwarmManager() *docker.Client {
	return client.swarmManager
}

// create a daemon docker client base on host id stored in ctx
func (client *CraneDockerClient) createNodeClient(nodeId string) (*docker.Client, error) {
	var swarmNode *docker.Client
	nodeUrl, err := client.NodeDaemonUrl(nodeId)
	if err != nil {
		return nil, err
	}

	endpoint := nodeUrl.String()
	if nodeUrl.Scheme == "https" {
		swarmNode, err = NewGoDockerClientTls(endpoint, API_VERSION, client.config)
	} else {
		swarmNode, err = docker.NewVersionedClient(endpoint, API_VERSION)
	}

	if err != nil {
		return nil, &cranerror.CraneError{
			Code: CodeConnToNodeError,
			Err:  &cranerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: err},
		}
	}

	return swarmNode, nil
}

// create node client: form manager node got endpoint by node label and verify node id
// by get docker info form httpclient
func (client *CraneDockerClient) SwarmNode(ctx context.Context) (*docker.Client, error) {
	nodeId, ok := ctx.Value("node_id").(string)
	if !ok {
		return nil, &cranerror.CraneError{
			Code: CodeConnToNodeError,
			Err: &cranerror.NodeConnError{
				ID:       nodeId,
				Endpoint: "",
				Err:      fmt.Errorf("parse node id  failed"),
			},
		}
	}

	nodeUrl, err := client.NodeDaemonUrl(nodeId)
	if err != nil {
		return nil, err
	}

	if err := client.VerifyNodeEndpoint(nodeId, nodeUrl); err != nil {
		return nil, err
	}

	nodeClient, err := client.createNodeClient(nodeId)
	if err != nil {
		return nil, err
	}

	return nodeClient, nil
}

func (client *CraneDockerClient) VerifyNodeEndpoint(nodeId string, nodeUrl *url.URL) error {
	if nodeUrl == nil {
		return &cranerror.CraneError{
			Code: CodeGetNodeEndpointError,
			Err:  &cranerror.NodeConnError{ID: nodeId, Err: fmt.Errorf("verify endpoint failed: empty node url")},
		}
	}
	var nodeInfo types.Info
	//TODO hardcode may have better way
	if nodeUrl.Scheme == "tcp" {
		nodeUrl.Scheme = "http"
	}

	endpoint := nodeUrl.String()
	content, err := client.sharedHttpClient.GET(nil, endpoint+"/info", url.Values{}, nil)
	if err != nil {
		return &cranerror.CraneError{
			Code: CodeVerifyNodeEnpointFailed,
			Err:  &cranerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: fmt.Errorf("verify endpoint failed: %s", err.Error())},
		}
	}

	if err := json.Unmarshal(content, &nodeInfo); err != nil {
		return &cranerror.CraneError{
			Code: CodeVerifyNodeEnpointFailed,
			Err:  &cranerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: fmt.Errorf("verify endpoint failed: %s", err.Error())},
		}
	}

	if err != nil {
		return &cranerror.CraneError{
			Code: CodeConnToNodeError,
			Err:  &cranerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: err},
		}
	}

	if nodeId != nodeInfo.Swarm.NodeID {
		return &cranerror.CraneError{
			Code: CodeNodeEndpointIpMatchError,
			Err:  &cranerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: fmt.Errorf("node id not matched endpoint")},
		}
	}

	return nil
}

func NewGoDockerClientTls(endpoint string, apiVersion string, config *config.Config) (*docker.Client, error) {
	tlsCaCert, tlsCert, tlsKey := SharedClientCertFiles(config)
	return docker.NewVersionedTLSClient(endpoint, tlsCert, tlsKey, tlsCaCert, apiVersion)
}

func NewHttpClient() (*httpclient.Client, error) {
	httpClient := &http.Client{Timeout: defaultHttpRequestTimeout}
	return httpclient.NewClient(httpClient, nil)
}

func NewHttpClientTls(config *config.Config) (*httpclient.Client, error) {
	tlsCaCert, tlsCert, tlsKey := SharedClientCertFiles(config)
	httpClient := &http.Client{Timeout: defaultHttpRequestTimeout}
	return httpclient.NewTLSClient(tlsCaCert, tlsCert, tlsKey, httpClient, nil)
}

func SharedClientCertFiles(config *config.Config) (string, string, string) {
	tlsCaCert := filepath.Join(config.DockerCertPath, "ca.pem")
	tlsCert := filepath.Join(config.DockerCertPath, "cert.pem")
	tlsKey := filepath.Join(config.DockerCertPath, "key.pem")

	return tlsCaCert, tlsCert, tlsKey
}

func ToCraneError(err error) error {
	var detailError error
	switch err.(type) {
	case *docker.NoSuchContainer:
		detailError = cranerror.NewError(CodeContainerInvalid, err.Error())
	case *docker.NoSuchNetwork:
		detailError = cranerror.NewError(CodeNetworkInvalid, err.Error())
	case *docker.NoSuchNetworkOrContainer:
		detailError = cranerror.NewError(CodeNetworkOrContainerInvalid, err.Error())
	case *docker.ContainerAlreadyRunning:
		detailError = cranerror.NewError(CodeContainerAlreadyRunning, err.Error())
	case *docker.ContainerNotRunning:
		detailError = cranerror.NewError(CodeContainerNotRunning, err.Error())
	default:
		detailError = err
	}

	return detailError
}
