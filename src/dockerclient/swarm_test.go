package dockerclient

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	mock "github.com/Dataman-Cloud/crane/src/testing"
	"github.com/Dataman-Cloud/crane/src/utils/config"
)

func TestPing(t *testing.T) {
	mockServer := mock.NewServer()
	defer mockServer.Close()

	envs := map[string]interface{}{
		"Version":       "1.10.1",
		"Os":            "linux",
		"KernelVersion": "3.13.0-77-generic",
		"GoVersion":     "go1.4.2",
		"GitCommit":     "9e83765",
		"Arch":          "amd64",
		"ApiVersion":    "1.22",
		"BuildTime":     "2015-12-01T07:09:13.444803460+00:00",
		"Experimental":  false,
	}
	mockServer.AddRouter("/_ping", "get").RGroup().
		Reply(200)
	mockServer.AddRouter("/version", "get").RGroup().
		Reply(200).
		WJSON(envs)

	mockServer.Register()

	config := &config.Config{
		DockerEntryScheme: mockServer.Scheme,
		SwarmManagerIP:    mockServer.Addr,
		DockerEntryPort:   mockServer.Port,
		DockerTlsVerify:   false,
		DockerApiVersion:  "",
	}
	craneDockerClient, err := NewCraneDockerClient(config)
	assert.Nil(t, err)

	err = craneDockerClient.Ping()
	assert.Nil(t, err)
}

// TODO (wtzhou) refactor me by assist package mock
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
