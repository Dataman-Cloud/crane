package dockerclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/Dataman-Cloud/rolex/src/utils/config"
	"github.com/Dataman-Cloud/rolex/src/utils/rolexerror"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	"github.com/Dataman-Cloud/rolex/src/utils/httpclient"
	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
)

type RolexDockerClient struct {
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

// initialize rolex docker client
func NewRolexDockerClient(config *config.Config) (*RolexDockerClient, error) {
	var err error

	swarmManagerEntry := config.DockerEntryScheme + "://" + config.SwarmManagerIP + ":" + config.DockerEntryPort
	client := &RolexDockerClient{
		config: config,

		swarmManagerHttpEndpoint: swarmManagerEntry,
	}

	if config.DockerTlsVerify {
		client.swarmManager, err = client.NewGoDockerClientTls(swarmManagerEntry, API_VERSION)
		client.sharedHttpClient, err = client.NewHttpClientTls()
	} else {
		client.swarmManager, err = docker.NewVersionedClient(swarmManagerEntry, API_VERSION)
		client.sharedHttpClient, err = client.NewHttpClient()
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
func (client *RolexDockerClient) SwarmManager() *docker.Client {
	return client.swarmManager
}

// create a daemon docker client base on host id stored in ctx
func (client *RolexDockerClient) createNodeClient(nodeId string) (*docker.Client, error) {
	var swarmNode *docker.Client
	nodeUrl, err := client.NodeDaemonUrl(nodeId)
	if err != nil {
		return nil, err
	}

	endpoint := nodeUrl.String()
	if nodeUrl.Scheme == "https" {
		swarmNode, err = client.NewGoDockerClientTls(endpoint, API_VERSION)
	} else {
		swarmNode, err = docker.NewVersionedClient(endpoint, API_VERSION)
	}

	if err != nil {
		return nil, &rolexerror.DmError{
			Code: CodeConnToNodeError,
			Err:  &rolexerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: err},
		}
	}

	return swarmNode, nil
}

// create node client: form manager node got endpoint by node label and verify node id
// by get docker info form httpclient
func (client *RolexDockerClient) SwarmNode(ctx context.Context) (*docker.Client, error) {
	nodeId, ok := ctx.Value("node_id").(string)
	if !ok {
		return nil, &rolexerror.DmError{
			Code: CodeConnToNodeError,
			Err: &rolexerror.NodeConnError{
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

func (client *RolexDockerClient) VerifyNodeEndpoint(nodeId string, nodeUrl *url.URL) error {
	if nodeUrl == nil {
		return &rolexerror.DmError{
			Code: CodeGetNodeEndpointError,
			Err:  &rolexerror.NodeConnError{ID: nodeId, Err: fmt.Errorf("verify endpoint failed: empty node url")},
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
		return &rolexerror.DmError{
			Code: CodeVerifyNodeEnpointFailed,
			Err:  &rolexerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: fmt.Errorf("verify endpoint failed: %s", err.Error())},
		}
	}

	if err := json.Unmarshal(content, &nodeInfo); err != nil {
		return &rolexerror.DmError{
			Code: CodeVerifyNodeEnpointFailed,
			Err:  &rolexerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: fmt.Errorf("verify endpoint failed: %s", err.Error())},
		}
	}

	if err != nil {
		return &rolexerror.DmError{
			Code: CodeConnToNodeError,
			Err:  &rolexerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: err},
		}
	}

	if nodeId != nodeInfo.Swarm.NodeID {
		return &rolexerror.DmError{
			Code: CodeNodeEndpointIpMatchError,
			Err:  &rolexerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: fmt.Errorf("node id not matched endpoint")},
		}
	}

	return nil
}

func (client *RolexDockerClient) NewGoDockerClientTls(endpoint string, apiVersion string) (*docker.Client, error) {
	tlsCaCert, tlsCert, tlsKey := SharedClientCertFiles(client.config)
	return docker.NewVersionedTLSClient(endpoint, tlsCert, tlsKey, tlsCaCert, apiVersion)
}

func (client *RolexDockerClient) NewHttpClient() (*httpclient.Client, error) {
	httpClient := &http.Client{Timeout: defaultHttpRequestTimeout}
	return httpclient.NewClient(httpClient, nil)
}

func (client *RolexDockerClient) NewHttpClientTls() (*httpclient.Client, error) {
	tlsCaCert, tlsCert, tlsKey := SharedClientCertFiles(client.config)
	httpClient := &http.Client{Timeout: defaultHttpRequestTimeout}
	return httpclient.NewTLSClient(tlsCaCert, tlsCert, tlsKey, httpClient, nil)
}

func SharedClientCertFiles(config *config.Config) (string, string, string) {
	tlsCaCert := filepath.Join(config.DockerCertPath, "ca.pem")
	tlsCert := filepath.Join(config.DockerCertPath, "cert.pem")
	tlsKey := filepath.Join(config.DockerCertPath, "key.pem")

	return tlsCaCert, tlsCert, tlsKey
}

func ToRolexError(err error) error {
	var detailError error
	switch err.(type) {
	case *docker.NoSuchContainer:
		detailError = rolexerror.NewError(CodeContainerNotFound, err.Error())
	case *docker.NoSuchNetwork:
		detailError = rolexerror.NewError(CodeNetworkNotFound, err.Error())
	case *docker.NoSuchNetworkOrContainer:
		detailError = rolexerror.NewError(CodeNetworkOrContainerNotFound, err.Error())
	case *docker.ContainerAlreadyRunning:
		detailError = rolexerror.NewError(CodeContainerAlreadyRunning, err.Error())
	case *docker.ContainerNotRunning:
		detailError = rolexerror.NewError(CodeContainerNotRunning, err.Error())
	default:
		detailError = err
	}

	return detailError
}
