package dockerclient

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/Dataman-Cloud/rolex/src/util/config"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"

	log "github.com/Sirupsen/logrus"
	goclient "github.com/fsouza/go-dockerclient"
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

	// client connect to swarm cluster
	swarmClient *goclient.Client
	// map clients connect to individual node
	dockerClients map[string]*goclient.Client
	// mutex control writing to dockerClients
	DockerClientMutex *sync.Mutex

	// http client shared both for cluster connection & client connection
	SharedHttpClient  *http.Client
	SwarmHttpEndpoint string

	Config *config.Config
}

// initialize rolex docker client
func NewRolexDockerClient(config *config.Config) (*RolexDockerClient, error) {
	var err error

	client := &RolexDockerClient{
		Config:            config,
		dockerClients:     make(map[string](*goclient.Client), 0),
		SwarmHttpEndpoint: strings.Replace(config.DockerHost, "tcp", "https", -1),
		DockerClientMutex: &sync.Mutex{},
	}

	if config.DockerTlsVerify == "1" {
		client.swarmClient, err = client.NewGoDockerClientTls(config.DockerHost, API_VERSION)
		client.SharedHttpClient, err = client.NewHttpClientTls()
	} else {
		client.swarmClient, err = goclient.NewVersionedClient(config.DockerHost, API_VERSION)
		client.SharedHttpClient = &http.Client{Timeout: defaultHttpRequestTimeout}
	}

	if err != nil {
		log.Error("Unable to connect to docker daemon . Ensure docker is running endpoint ", config.DockerHost, "err: ", err)
		return nil, err
	}

	err = client.swarmClient.Ping()
	if err != nil {
		log.Error("Unable to ping docker daemon. Ensure docker is running endpoint ", config.DockerHost, "err: ", err)
		return nil, err
	}

	return client, nil
}

// return swarm docker client
func (client *RolexDockerClient) SwarmClient() *goclient.Client {
	return client.swarmClient
}

// return or cache daemon docker client base on host id stored in ctx
func (client *RolexDockerClient) DockerClient(ctx context.Context) (*goclient.Client, error) {
	var dockerClient *goclient.Client

	var err error
	node_id, ok := ctx.Value("node_id").(string)
	if !ok {
		err = rolexerror.NewRolexError(rolexerror.CodeGetDockerClientError, "node_id not found")
		return nil, err
	}

	if dockerClient, found := client.dockerClients[node_id]; found {
		return dockerClient, nil
	}

	host, err := client.NodeDaemonEndpoint(node_id, "tcp")
	if err != nil {
		err = rolexerror.NewRolexError(rolexerror.CodeGetDockerClientError, "unable to parse node ip for "+host)
		return nil, err
	}

	if client.Config.DockerTlsVerify == "1" {
		dockerClient, err = client.NewGoDockerClientTls(host, API_VERSION)
	} else {
		dockerClient, err = goclient.NewVersionedClient(host, API_VERSION)
	}

	if err != nil {
		message := fmt.Sprintf("failed to init client %s error: %s", host, err.Error())
		err = rolexerror.NewRolexError(rolexerror.CodeGetDockerClientError, message)
		return nil, err
	}

	err = dockerClient.Ping()
	if err != nil {
		message := fmt.Sprintf("DockerClient ping error: %s", err.Error())
		err = rolexerror.NewRolexError(rolexerror.CodeGetDockerClientError, message)
		return nil, err
	}

	client.DockerClientMutex.Lock()
	defer client.DockerClientMutex.Unlock()
	client.dockerClients[node_id] = dockerClient

	return dockerClient, nil
}

// ping to test swarmClient connection
func (client *RolexDockerClient) Ping() error {
	return client.swarmClient.Ping()
}

func (client *RolexDockerClient) NewGoDockerClientTls(endpoint string, apiVersion string) (*goclient.Client, error) {
	tlsCaCert, tlsCert, tlsKey := SharedClientCertFiles(client.Config)
	return goclient.NewVersionedTLSClient(endpoint, tlsCert, tlsKey, tlsCaCert, apiVersion)
}

func (client *RolexDockerClient) NewHttpClientTls() (*http.Client, error) {
	tlsCaCert, tlsCert, tlsKey := SharedClientCertFiles(client.Config)

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

func SortingError(err error) error {
	var detailError error
	switch err.(type) {
	case *goclient.NoSuchContainer:
		detailError = rolexerror.NewRolexError(rolexerror.CodeContainerNotFound, err.Error())
	case *goclient.NoSuchNetwork:
		detailError = rolexerror.NewRolexError(rolexerror.CodeNetworkNotFound, err.Error())
	case *goclient.NoSuchNetworkOrContainer:
		detailError = rolexerror.NewRolexError(rolexerror.CodeNetworkOrContainerNotFound, err.Error())
	case *goclient.ContainerAlreadyRunning:
		detailError = rolexerror.NewRolexError(rolexerror.CodeContainerAlreadyRunning, err.Error())
	case *goclient.ContainerNotRunning:
		detailError = rolexerror.NewRolexError(rolexerror.CodeContainerNotRunning, err.Error())
	default:
		detailError = err
	}

	return detailError
}
