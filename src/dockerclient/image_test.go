package dockerclient

import (
	"testing"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestImage(t *testing.T) {
	testServer, craneClient, nodeId := InitTestSwarm(t)
	assert.NotNil(t, nodeId)
	assert.NotNil(t, craneClient)
	defer testServer.Stop()
	backgroundContext := context.Background()
	craneContext := context.WithValue(backgroundContext, "node_id", "errorid")
	_, err := craneClient.ListImages(craneContext, docker.ListImagesOptions{})
	assert.NotNil(t, err)
	_, err = craneClient.InspectImage(craneContext, "test")
	assert.NotNil(t, err)
	_, err = craneClient.ImageHistory(craneContext, "test")
	assert.NotNil(t, err)
	err = craneClient.RemoveImage(craneContext, "test")
	assert.NotNil(t, err)

	endpoint := testServer.URL()[0 : len(testServer.URL())-1]
	_, err = craneClient.sharedHttpClient.POST(nil, endpoint+"/build", nil, nil, nil)
	assert.Nil(t, err)

	craneContext = context.WithValue(backgroundContext, "node_id", nodeId)
	images, err := craneClient.ListImages(craneContext, docker.ListImagesOptions{All: true})
	assert.Nil(t, err)
	assert.NotNil(t, images)
	assert.Equal(t, len(images), 1)

	imageId := images[0].ID
	image, err := craneClient.InspectImage(craneContext, imageId)
	assert.Nil(t, err)
	assert.NotNil(t, image)
	assert.Equal(t, imageId, image.ID)

	//TODO: (upccup) add image history to fake server
	//imageHistory, err := craneClient.ImageHistory(craneContext, imageId)
	//assert.Nil(t, err)
	//assert.NotNil(t, imageHistory)

	err = craneClient.RemoveImage(craneContext, imageId)
	assert.Nil(t, err)
}
