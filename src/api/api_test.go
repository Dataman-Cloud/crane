package api

import (
	"testing"

	"github.com/Dataman-Cloud/crane/src/dockerclient"
	"github.com/Dataman-Cloud/crane/src/utils/config"

	"github.com/stretchr/testify/assert"
)

func TestGetDockerClient(t *testing.T) {
	fakeClient := &dockerclient.CraneDockerClient{}

	fakeApi := &Api{
		Client: fakeClient,
	}

	client := fakeApi.GetDockerClient()
	assert.Equal(t, fakeClient, client)
}

func TestGetConfig(t *testing.T) {
	fakeConfig := &config.Config{}

	fakeApi := &Api{
		Config: fakeConfig,
	}

	config := fakeApi.GetConfig()
	assert.Equal(t, fakeConfig, config)
}
