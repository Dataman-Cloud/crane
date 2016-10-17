package dockerclient

import (
	"os"
	"testing"

	"github.com/Dataman-Cloud/crane/src/utils/cranerror"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestVolume(t *testing.T) {
	os.Setenv("CRANE_ADDR", "foobar")
	os.Setenv("CRANE_SWARM_MANAGER_IP", "foobar")
	os.Setenv("CRANE_DOCKER_CERT_PATH", "foobar")
	os.Setenv("CRANE_DB_DRIVER", "foobar")
	os.Setenv("CRANE_DB_DSN", "foobar")
	os.Setenv("CRANE_FEATURE_FLAGS", "foobar")
	os.Setenv("CRANE_REGISTRY_PRIVATE_KEY_PATH", "foobar")
	os.Setenv("CRANE_REGISTRY_ADDR", "foobar")
	os.Setenv("CRANE_ACCOUNT_AUTHENTICATOR", "foobar")
	defer os.Setenv("CRANE_ADDR", "")
	defer os.Setenv("CRANE_SWARM_MANAGER_IP", "")
	defer os.Setenv("CRANE_DOCKER_CERT_PATH", "")
	defer os.Setenv("CRANE_DB_DRIVER", "")
	defer os.Setenv("CRANE_DB_DSN", "")
	defer os.Setenv("CRANE_FEATURE_FLAGS", "")
	defer os.Setenv("CRANE_REGISTRY_PRIVATE_KEY_PATH", "")
	defer os.Setenv("CRANE_REGISTRY_ADDR", "")
	defer os.Setenv("CRANE_ACCOUNT_AUTHENTICATOR", "")

	testServer, craneClient, nodeId := InitTestSwarm(t)
	assert.NotNil(t, nodeId)
	assert.NotNil(t, craneClient)
	defer testServer.Stop()
	backgroundContext := context.Background()
	craneContext := context.WithValue(backgroundContext, "node_id", "errorid")
	_, err := craneClient.CreateVolume(craneContext, docker.CreateVolumeOptions{})
	assert.NotNil(t, err)
	_, err = craneClient.InspectVolume(craneContext, "test")
	assert.NotNil(t, err)
	_, err = craneClient.ListVolumes(craneContext, docker.ListVolumesOptions{})
	assert.NotNil(t, err)
	err = craneClient.RemoveVolume(craneContext, "test")
	assert.NotNil(t, err)

	endpoint := testServer.URL()[0 : len(testServer.URL())-1]
	_, err = craneClient.sharedHttpClient.POST(nil, endpoint+"/build", nil, nil, nil)
	assert.Nil(t, err)

	craneContext = context.WithValue(backgroundContext, "node_id", nodeId)
	volume, err := craneClient.CreateVolume(craneContext, docker.CreateVolumeOptions{Name: "^^^^adsasda!@"})
	assert.Nil(t, volume)
	assert.NotNil(t, err)
	craneErr, ok := err.(*cranerror.CraneError)
	assert.True(t, ok)
	assert.Equal(t, craneErr.Code, CodeInvalidVolumeName)

	volume, err = craneClient.CreateVolume(craneContext, docker.CreateVolumeOptions{Name: "testupc"})
	assert.Nil(t, err)
	assert.NotNil(t, volume)
	assert.Equal(t, volume.Name, "testupc")

	volumes, err := craneClient.ListVolumes(craneContext, docker.ListVolumesOptions{})
	assert.Nil(t, err)
	assert.NotNil(t, volumes)
	assert.Equal(t, len(volumes), 1)

	volume, err = craneClient.InspectVolume(craneContext, "testupc")
	assert.Nil(t, err)
	assert.NotNil(t, volume)
	assert.Equal(t, volume.Name, "testupc")

	err = craneClient.RemoveVolume(craneContext, "testupc")
	assert.Nil(t, err)
}
