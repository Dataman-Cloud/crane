package dockerclient

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/docker/engine-api/types/swarm"
	"github.com/stretchr/testify/assert"
)

func TestToCraneServiceSpec(t *testing.T) {
	body := `
	{
	        "Name": "none",
	        "Id": "1836d62be355e36050913f118835bd1fd6be10638e799ccaf5ea76bc6820ced2",
	        "Scope": "local",
	        "Driver": "null",
	        "EnableIPv6": false,
	        "IPAM": {
                        "Driver": "default",
	                "Options": null,
	                "Config": []
                 },
                "Internal": false,
                "Containers": {},
                "Options": {},
	        "Labels": {}
	 }`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(body))
	}))

	defer server.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)

	client := &CraneDockerClient{
		sharedHttpClient: httpClient,
	}

	craneServiceSpe := client.ToCraneServiceSpec(swarm.ServiceSpec{})
	assert.NotNil(t, craneServiceSpe)
}

func TestParseEndpoint(t *testing.T) {
	_, err := parseEndpoint("localhost:2375")
	assert.Nil(t, err)

	_, err = parseEndpoint("tcp://localhost:2375")
	assert.Nil(t, err)

	_, err = parseEndpoint("tcp://localhost")
	assert.Nil(t, err)

	_, err = parseEndpoint("://localhost:2375")
	assert.NotNil(t, err)
}

func TestgetAdvertiseAddrByEndpoint(t *testing.T) {
	var host string
	var err error
	host, err = getAdvertiseAddrByEndpoint("localhost:2375")
	assert.Nil(t, err)
	assert.Equal(t, host, "localhost")
	host, err = getAdvertiseAddrByEndpoint("tcp://localhost:2375")
	assert.Nil(t, err)
	assert.Equal(t, host, "localhost")
	host, err = getAdvertiseAddrByEndpoint("localhost")
	assert.Nil(t, err)
	assert.Equal(t, host, "localhost")
	host, err = getAdvertiseAddrByEndpoint("http://localhost")
	assert.Nil(t, err)
	assert.Equal(t, host, "localhost")
	host, err = getAdvertiseAddrByEndpoint("://localhost")
	assert.NotNil(t, err)
}

func TestGetServicesNamespace(t *testing.T) {
	namespace := GetServicesNamespace(swarm.ServiceSpec{})
	assert.Equal(t, namespace, "")

	spec := swarm.ServiceSpec{}
	spec.Annotations.Labels = map[string]string{
		labelNamespace: "value",
	}
	namespace = GetServicesNamespace(spec)
	assert.Equal(t, namespace, "value")
}
