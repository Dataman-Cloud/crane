package dockerclient

import (
	"encoding/json"
	"testing"

	"github.com/Dataman-Cloud/crane/src/model"

	dockertest "github.com/Dataman-Cloud/go-dockerclient/testing"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestListImage(t *testing.T) {
	_, err := dockertest.NewServer("127.0.0.1:8888", nil, nil)
	assert.Nil(t, err)

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)
	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: "http://127.0.0.1:8888",
	}

	swarmInitInfo := swarm.InitRequest{
		//ListenAddr:      "0.0.0.0:2377",
		//AdvertiseAddr:   "192.168.1.1:2377",
		ForceNewCluster: false,
	}
	_, err = client.sharedHttpClient.POST(nil, client.swarmManagerHttpEndpoint+"/swarm/init", nil, swarmInitInfo, nil)
	assert.Nil(t, err)
	nodes, err := client.ListNode(types.NodeListOptions{})
	assert.Equal(t, len(nodes), 1)
	node := nodes[0]
	var nodeUpdate model.UpdateOptions
	updateOptions := `{"Method":"endpoint-update", "Options": "http://127.0.0.1:8888"}`
	err = json.Unmarshal([]byte(updateOptions), &nodeUpdate)
	err = client.UpdateNode(node.ID, nodeUpdate)
	assert.Nil(t, err)
	nodeUrl, err := client.GetDaemonUrlById(node.ID)
	assert.Nil(t, err)

	backgroundContext := context.Background()
	craneContext := context.WithValue(backgroundContext, "node_id", node.ID)
	dockerClient, err := client.SwarmNode(craneContext)
	assert.NotNil(t, dockerClient)
	t.Log(nodeUrl)
}
