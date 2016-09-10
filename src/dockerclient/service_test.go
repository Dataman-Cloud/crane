package dockerclient

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/stretchr/testify/assert"
)

func TestCreateService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"ID":"e90302"}`))
	}))
	defer server.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)

	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server.URL,
	}

	_, err = client.CreateService(swarm.ServiceSpec{}, types.ServiceCreateOptions{})
	assert.Nil(t, err)

	client = &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: "errorurl",
	}

	_, err = client.CreateService(swarm.ServiceSpec{}, types.ServiceCreateOptions{})
	assert.NotNil(t, err)

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`error response`))
	}))
	client = &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server.URL,
	}

	_, err = client.CreateService(swarm.ServiceSpec{}, types.ServiceCreateOptions{})
	assert.NotNil(t, err)
}

func TestListServiceSpec(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[{"ID":"e90302"}]`))
	}))
	defer server.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)
	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server.URL,
	}
	_, err = client.ListServiceSpec(types.ServiceListOptions{})
	assert.Nil(t, err)

	client = &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: "error url",
	}
	_, err = client.ListServiceSpec(types.ServiceListOptions{})
	assert.NotNil(t, err)

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`error response`))
	}))
	client = &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server.URL,
	}
	_, err = client.ListServiceSpec(types.ServiceListOptions{})
	assert.NotNil(t, err)
}

func TestListService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[{"ID":"e90302"}]`))
	}))
	defer server.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)
	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server.URL,
	}
	_, err = client.ListService(types.ServiceListOptions{})
	assert.Nil(t, err)

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`error response`))
	}))
	client = &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: "testurl",
	}
	_, err = client.ListService(types.ServiceListOptions{})
	assert.NotNil(t, err)
}
