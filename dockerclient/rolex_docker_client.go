package dockerclient

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
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
	DockerClient *goclient.Client
	HttpClient   *http.Client
	HttpEndpoint string
}

func (dg *RolexDockerClient) Ping() error {
	return dg.DockerClient.Ping()
}

func NewRolexDockerClient(config *config.Config) (*RolexDockerClient, error) {
	var err error
	var client *goclient.Client
	var httpClient *http.Client

	if config.DockerTlsVerify == "1" {
		tlsCaCert := filepath.Join(config.DockerCertPath, "ca.pem")
		tlsCert := filepath.Join(config.DockerCertPath, "cert.pem")
		tlsKey := filepath.Join(config.DockerCertPath, "key.pem")
		client, err = goclient.NewVersionedTLSClient(config.DockerHost, tlsCert, tlsKey, tlsCaCert, "1.23")

		// Load CA cert
		caCert, err := ioutil.ReadFile(tlsCaCert)
		if err != nil {
			log.Fatal(err)
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

		httpClient = &http.Client{Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
			Timeout: defaultHttpRequestTimeout,
		}
	} else {
		client, err = goclient.NewVersionedClient(config.DockerHost, "1.23")
		httpClient = &http.Client{Timeout: defaultHttpRequestTimeout}
	}

	if err != nil {
		log.Error("Unable to connect to docker daemon . Ensure docker is running endpoint ", config.DockerHost, "err", err)
		return nil, err
	}

	err = client.Ping()
	if err != nil {
		log.Error("Unable to ping docker daemon. Ensure docker is running endpoint ", config.DockerHost, "err", err)
		return nil, err
	}

	return &RolexDockerClient{
		DockerClient: client,
		HttpClient:   httpClient,
		HttpEndpoint: strings.Replace(config.DockerHost, "tcp", "https", -1),
	}, nil
}

// execute http get request use default timeout
func (client *RolexDockerClient) HttpGet(requestPath string) ([]byte, error) {
	resp, err := client.HttpClient.Get(client.HttpEndpoint + "/" + requestPath)
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
		return nil, fmt.Errorf("http response status code is %d not 200", resp.StatusCode)
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
		return nil, fmt.Errorf("http response status code is %d not 200", resp.StatusCode)
	}

	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)

}
