package dockerclient

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/Dataman-Cloud/rolex/util/config"

	log "github.com/Sirupsen/logrus"
	goclient "github.com/fsouza/go-dockerclient"
)

const (
	defaultHttpRequestTimeout = time.Second * 10
)

type RolexDockerClient struct {
	DockerClientInterface
	swarmClient  *goclient.Client // client connect to swarm cluster
	dockerClient *goclient.Client // client connect to individual node
	HttpClient   *http.Client
	HttpEndpoint string
}

func (client *RolexDockerClient) SwarmClient() *goclient.Client {
	return client.swarmClient
}

func (client *RolexDockerClient) DockerClient(nodeId string) *goclient.Client {
	//TODO connect to a docker node
	return client.swarmClient
}

func (client *RolexDockerClient) Ping() error {
	return client.SwarmClient().Ping()
}

func NewRolexDockerClient(config *config.Config) (*RolexDockerClient, error) {
	var err error
	var swarmClient *goclient.Client
	var httpClient *http.Client

	if config.DockerTlsVerify == "1" {
		swarmClient, err = NewSwarmClientTls(config)
		httpClient, err = NewHttpClientTls(config)
	} else {
		swarmClient, err = goclient.NewVersionedClient(config.DockerHost, "1.23")
		httpClient = &http.Client{Timeout: defaultHttpRequestTimeout}
	}

	if err != nil {
		log.Error("Unable to connect to docker daemon . Ensure docker is running endpoint ", config.DockerHost, "err: ", err)
		return nil, err
	}

	err = swarmClient.Ping()
	if err != nil {
		log.Error("Unable to ping docker daemon. Ensure docker is running endpoint ", config.DockerHost, "err: ", err)
		return nil, err
	}

	return &RolexDockerClient{
		swarmClient:  swarmClient,
		HttpClient:   httpClient,
		HttpEndpoint: strings.Replace(config.DockerHost, "tcp", "https", -1),
	}, nil
}

func NewSwarmClientTls(config *config.Config) (*goclient.Client, error) {
	tlsCaCert := filepath.Join(config.DockerCertPath, "ca.pem")
	tlsCert := filepath.Join(config.DockerCertPath, "cert.pem")
	tlsKey := filepath.Join(config.DockerCertPath, "key.pem")
	return goclient.NewVersionedTLSClient(config.DockerHost, tlsCert, tlsKey, tlsCaCert, "1.23")
}

func NewHttpClientTls(config *config.Config) (*http.Client, error) {
	tlsCaCert := filepath.Join(config.DockerCertPath, "ca.pem")
	tlsCert := filepath.Join(config.DockerCertPath, "cert.pem")
	tlsKey := filepath.Join(config.DockerCertPath, "key.pem")

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
