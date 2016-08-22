package dockerclient

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/Dataman-Cloud/rolex/src/util/config"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"

	"github.com/Dataman-Cloud/go-component/utils/httpclient"
	docker "github.com/Dataman-Cloud/go-dockerclient"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
)

const (
	defaultHttpRequestTimeout = time.Second * 10
)

const (
	API_VERSION = "1.23"
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

// return or cache daemon docker client base on host id stored in ctx
func (client *RolexDockerClient) SwarmNode(ctx context.Context) (*docker.Client, error) {
	var swarmNode *docker.Client
	var err error
	nodeId, ok := ctx.Value("node_id").(string)
	if !ok {
		return nil, &rolexerror.RolexError{
			Code: rolexerror.CodeConnToNodeError,
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

	endpoint := nodeUrl.String()
	if nodeUrl.Scheme == "https" {
		swarmNode, err = client.NewGoDockerClientTls(endpoint, API_VERSION)
	} else {
		swarmNode, err = docker.NewVersionedClient(endpoint, API_VERSION)
	}

	if err != nil {
		return nil, &rolexerror.RolexError{
			Code: rolexerror.CodeConnToNodeError,
			Err:  &rolexerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: err},
		}
	}

	nodeInfo, err := client.Info(nodeId)
	if err != nil {
		return nil, &rolexerror.RolexError{
			Code: rolexerror.CodeConnToNodeError,
			Err:  &rolexerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: err},
		}
	}

	if nodeId != nodeInfo.Swarm.NodeID {
		return nil, &rolexerror.RolexError{
			Code: rolexerror.CodeNodeEndpointIpMatchError,
			Err:  &rolexerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: fmt.Errorf("node id not matched endpoint")},
		}
	}
	return swarmNode, nil
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
		detailError = rolexerror.NewRolexError(rolexerror.CodeContainerNotFound, err.Error())
	case *docker.NoSuchNetwork:
		detailError = rolexerror.NewRolexError(rolexerror.CodeNetworkNotFound, err.Error())
	case *docker.NoSuchNetworkOrContainer:
		detailError = rolexerror.NewRolexError(rolexerror.CodeNetworkOrContainerNotFound, err.Error())
	case *docker.ContainerAlreadyRunning:
		detailError = rolexerror.NewRolexError(rolexerror.CodeContainerAlreadyRunning, err.Error())
	case *docker.ContainerNotRunning:
		detailError = rolexerror.NewRolexError(rolexerror.CodeContainerNotRunning, err.Error())
	default:
		detailError = err
	}

	return detailError
}
