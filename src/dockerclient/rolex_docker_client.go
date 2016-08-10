package dockerclient

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/Dataman-Cloud/rolex/src/util/config"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"

	log "github.com/Sirupsen/logrus"
	docker "github.com/fsouza/go-dockerclient"
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
	// map clients connect to individual node
	swarmNodes map[string]*docker.Client
	// mutex control writing to swarmNodes
	swarmNodesMutex *sync.Mutex

	// http client shared both for cluster connection & client connection
	sharedHttpClient         *http.Client
	swarmManagerHttpEndpoint string
	swarmNodeHttpEndpoints   []string

	config *config.Config
}

// initialize rolex docker client
func NewRolexDockerClient(config *config.Config) (*RolexDockerClient, error) {
	var err error

	swarmManagerEntry := config.DockerEntryScheme + "://" + config.SwarmManagerIP + ":" + config.DockerEntryPort
	client := &RolexDockerClient{
		config: config,

		swarmNodes:      make(map[string](*docker.Client), 0),
		swarmNodesMutex: &sync.Mutex{},

		swarmManagerHttpEndpoint: swarmManagerEntry,
	}

	if config.DockerTlsVerify {
		client.swarmManager, err = client.NewGoDockerClientTls(swarmManagerEntry, API_VERSION)
		client.sharedHttpClient, err = client.NewHttpClientTls()
	} else {
		client.swarmManager, err = docker.NewVersionedClient(swarmManagerEntry, API_VERSION)
		client.sharedHttpClient = &http.Client{Timeout: defaultHttpRequestTimeout}
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

	client.swarmNodesMutex.Lock()
	defer client.swarmNodesMutex.Unlock()

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

	if swarmNode, found := client.swarmNodes[nodeId]; found {
		return swarmNode, nil
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

	err = swarmNode.Ping()
	if err != nil {
		return nil, &rolexerror.RolexError{
			Code: rolexerror.CodeConnToNodeError,
			Err:  &rolexerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: err},
		}
	}

	client.swarmNodes[nodeId] = swarmNode

	return swarmNode, nil
}

// ping to test swarmManager connection
func (client *RolexDockerClient) Ping() error {
	return client.swarmManager.Ping()
}

func (client *RolexDockerClient) NewGoDockerClientTls(endpoint string, apiVersion string) (*docker.Client, error) {
	tlsCaCert, tlsCert, tlsKey := SharedClientCertFiles(client.config)
	return docker.NewVersionedTLSClient(endpoint, tlsCert, tlsKey, tlsCaCert, apiVersion)
}

func (client *RolexDockerClient) NewHttpClientTls() (*http.Client, error) {
	tlsCaCert, tlsCert, tlsKey := SharedClientCertFiles(client.config)

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

	httpClient := &http.Client{Transport: &http.Transport{
		TLSClientConfig: tlsConfig,
	}, Timeout: defaultHttpRequestTimeout}

	return httpClient, nil
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
