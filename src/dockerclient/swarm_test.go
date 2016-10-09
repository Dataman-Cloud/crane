package dockerclient

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInspectSwarmErrorGet(t *testing.T) {
	fakeCluster := func(ctx *gin.Context) {
		var body swarm.Swarm
		body.JoinTokens.Manager = "FakeManagerToken"
		ctx.JSON(http.StatusBadRequest, body)
	}

	managerRouter := gin.New()
	managerRouter.GET("/swarm", fakeCluster)

	manager := httptest.NewServer(managerRouter)
	defer manager.Close()

	httpClient, err := NewHttpClient()
	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: manager.URL,
	}

	_, err = client.InspectSwarm()
	assert.NotNil(t, err)
}

func TestInspectSwarmErrorJSON(t *testing.T) {
	fakeCluster := func(ctx *gin.Context) {
		body := "fake body: cannot be json Unmarshal"
		ctx.JSON(http.StatusOK, body)
	}

	managerRouter := gin.New()
	managerRouter.GET("/swarm", fakeCluster)

	manager := httptest.NewServer(managerRouter)
	defer manager.Close()

	httpClient, err := NewHttpClient()
	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: manager.URL,
	}

	_, err = client.InspectSwarm()
	assert.NotNil(t, err)
}

func TestManagerInfoErrorGet(t *testing.T) {
	fakeNodeInfo := func(addr string, nodeID string) func(ctx *gin.Context) {
		nodeInfo := func(ctx *gin.Context) {
			body := types.Info{
				Swarm: swarm.Info{
					NodeAddr: addr,
					NodeID:   nodeID,
				},
			}
			ctx.JSON(http.StatusBadRequest, body)
		}
		return nodeInfo
	}

	managerRouter := gin.New()
	managerRouter.GET("/info", fakeNodeInfo("FakeAddr", "fakeManagerID"))

	manager := httptest.NewServer(managerRouter)
	defer manager.Close()

	httpClient, err := NewHttpClient()
	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: manager.URL,
	}

	_, err = client.ManagerInfo()
	assert.NotNil(t, err)
}

func TestManagerInfoErrorJSON(t *testing.T) {
	fakeNodeInfo := func(addr string, nodeID string) func(ctx *gin.Context) {
		nodeInfo := func(ctx *gin.Context) {
			body := "fake body: cannot be json Unmarshal"
			ctx.JSON(http.StatusOK, body)
		}
		return nodeInfo
	}

	managerRouter := gin.New()
	managerRouter.GET("/info", fakeNodeInfo("FakeAddr", "fakeManagerID"))

	manager := httptest.NewServer(managerRouter)
	defer manager.Close()

	httpClient, err := NewHttpClient()
	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: manager.URL,
	}

	_, err = client.ManagerInfo()
	assert.NotNil(t, err)
}
