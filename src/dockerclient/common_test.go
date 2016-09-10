package dockerclient

import (
	"testing"

	"github.com/docker/engine-api/types/swarm"
	"github.com/stretchr/testify/assert"
)

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
