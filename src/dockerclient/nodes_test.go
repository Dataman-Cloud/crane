package dockerclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/Dataman-Cloud/crane/src/model"
	mock "github.com/Dataman-Cloud/crane/src/testing"
	"github.com/Dataman-Cloud/crane/src/utils/config"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInspectNodeError(t *testing.T) {
	body := `{"Id":"e90302"}`
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(body))
	}))
	defer server1.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)

	client := &CraneDockerClient{
		sharedHttpClient: httpClient,
	}

	_, err = client.InspectNode("test")
	assert.NotNil(t, err)
}

func TestInspectNode(t *testing.T) {
	body := `
	{
	    "ID":"1t6jojzasio4veexyubvic4j2",
	    "Version":{
	        "Index":26607
	    },
	    "CreatedAt":"2016-08-26T08:00:24.466491891Z",
	    "UpdatedAt":"2016-09-08T05:23:49.697933079Z",
	    "Spec":{
	        "Labels":{
	            "dm.reserved.node.endpoint":"http://192.168.59.103:2376"
	        },
	        "Role":"worker",
	        "Availability":"active"
	    },
	    "Description":{
	        "Hostname":"192.168.59.013",
	        "Platform":{
	            "Architecture":"x86_64",
	            "OS":"linux"
	        },
	        "Resources":{
	            "NanoCPUs":2000000000,
	            "MemoryBytes":3975561216
	        },
	        "Engine":{
	            "EngineVersion":"1.12.0",
	            "Plugins":[
	                {
	                    "Type":"Network",
	                    "Name":"bridge"
	                },
	                {
	                    "Type":"Network",
	                    "Name":"host"
	                },
	                {
	                    "Type":"Network",
	                    "Name":"null"
	                },
	                {
	                    "Type":"Network",
	                    "Name":"overlay"
	                },
	                {
	                    "Type":"Volume",
	                    "Name":"local"
	                }
	            ]
	        }
	    },
	    "Status":{
	        "State":"down"
	    }
	}
	`
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(body))
	}))
	defer server1.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)

	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server1.URL,
	}

	_, err = client.InspectNode("test")
	assert.Nil(t, err)
}

func TestUpdateNodeRole(t *testing.T) {
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
	node := swarm.Node{
		ID: "nodeid",
		Spec: swarm.NodeSpec{
			Role: "manager",
		},
		Meta: swarm.Meta{
			Version: swarm.Version{
				Index: uint64(100),
			},
		},
	}
	opts := model.UpdateOptions{
		Method:  flagUpdateRole,
		Options: []byte(`"worker"`),
	}
	optsJSONError := opts
	optsJSONError.Options = []byte(`error json}`)

	requestBody := node.Spec
	requestBody.Role = "worker"

	mockServer.AddRouter("/_ping", "get").RGroup().
		Reply(200)
	mockServer.AddRouter("/version", "get").RGroup().
		Reply(200).
		WJSON(envs)
	mockServer.AddRouter("/nodes/nodeid/update", "post").RGroup().
		Reply(200).
		RQuery("version=100").
		RJSON(requestBody)

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

	err = craneDockerClient.UpdateNode(node, opts)
	assert.Nil(t, err)

	err = craneDockerClient.UpdateNode(node, optsJSONError)
	assert.NotNil(t, err)
}

func TestUpdateNodeAvailability(t *testing.T) {
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
	node := swarm.Node{
		ID: "nodeid",
		Spec: swarm.NodeSpec{
			Availability: "active",
		},
		Meta: swarm.Meta{
			Version: swarm.Version{
				Index: uint64(100),
			},
		},
	}
	opts := model.UpdateOptions{
		Method:  flagUpdateAvailability,
		Options: []byte(`"drain"`),
	}
	optsJSONError := opts
	optsJSONError.Options = []byte(`error json}`)

	requestBody := node.Spec
	requestBody.Availability = "drain"

	mockServer.AddRouter("/_ping", "get").RGroup().
		Reply(200)
	mockServer.AddRouter("/version", "get").RGroup().
		Reply(200).
		WJSON(envs)
	mockServer.AddRouter("/nodes/nodeid/update", "post").RGroup().
		Reply(200).
		RQuery("version=100").
		RJSON(requestBody)

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

	err = craneDockerClient.UpdateNode(node, opts)
	assert.Nil(t, err)
	err = craneDockerClient.UpdateNode(node, optsJSONError)
	assert.NotNil(t, err)
}

func TestUpdateNodeLabelAdd(t *testing.T) {
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
	node := swarm.Node{
		ID: "nodeid",
		Spec: swarm.NodeSpec{
			Annotations: swarm.Annotations{
				Labels: nil,
			},
		},
		Meta: swarm.Meta{
			Version: swarm.Version{
				Index: uint64(100),
			},
		},
	}
	opts := model.UpdateOptions{
		Method: flagLabelAdd,
		Options: []byte(`
		{
			"labelKey": "labelValue"
		}
		`),
	}
	optsJSONError := opts
	optsJSONError.Options = []byte(`error json}`)

	requestBody := node.Spec
	requestBody.Annotations.Labels = map[string]string{
		"labelKey": "labelValue",
	}

	mockServer.AddRouter("/_ping", "get").RGroup().
		Reply(200)
	mockServer.AddRouter("/version", "get").RGroup().
		Reply(200).
		WJSON(envs)
	mockServer.AddRouter("/nodes/nodeid/update", "post").RGroup().
		Reply(200).
		RQuery("version=100").
		RJSON(requestBody)

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

	err = craneDockerClient.UpdateNode(node, opts)
	assert.Nil(t, err)
	err = craneDockerClient.UpdateNode(node, optsJSONError)
	assert.NotNil(t, err)
}

func TestUpdateNodeLabelRemove(t *testing.T) {
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
	node := swarm.Node{
		ID: "nodeid",
		Spec: swarm.NodeSpec{
			Annotations: swarm.Annotations{
				Labels: map[string]string{
					"keyToBeRemoved": "valueToBeRemoved",
				},
			},
		},
		Meta: swarm.Meta{
			Version: swarm.Version{
				Index: uint64(100),
			},
		},
	}
	opts := model.UpdateOptions{
		Method:  flagLabelRemove,
		Options: []byte(`["keyToBeRemoved"]`),
	}
	optsJSONError := opts
	optsJSONError.Options = []byte(`error json}`)

	requestBody := node.Spec
	requestBody.Annotations.Labels = map[string]string{}

	mockServer.AddRouter("/_ping", "get").RGroup().
		Reply(200)
	mockServer.AddRouter("/version", "get").RGroup().
		Reply(200).
		WJSON(envs)
	mockServer.AddRouter("/nodes/nodeid/update", "post").RGroup().
		Reply(200).
		RQuery("version=100").
		RJSON(requestBody)

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

	err = craneDockerClient.UpdateNode(node, opts)
	assert.Nil(t, err)
	err = craneDockerClient.UpdateNode(node, optsJSONError)
	assert.NotNil(t, err)
}

func TestUpdateNodeLabelUpdate(t *testing.T) {
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
	node := swarm.Node{
		ID: "nodeid",
		Spec: swarm.NodeSpec{
			Annotations: swarm.Annotations{
				Labels: map[string]string{
					"keyToBeUpdated": "valueCurrent",
				},
			},
		},
		Meta: swarm.Meta{
			Version: swarm.Version{
				Index: uint64(100),
			},
		},
	}
	opts := model.UpdateOptions{
		Method:  flagLabelUpdate,
		Options: []byte(`{"keyToBeUpdated": "valueUpdated"}`),
	}
	optsJSONError := opts
	optsJSONError.Options = []byte(`error json}`)

	requestBody := node.Spec
	requestBody.Annotations.Labels["keyToBeUpdated"] = "valueUpdated"

	mockServer.AddRouter("/_ping", "get").RGroup().
		Reply(200)
	mockServer.AddRouter("/version", "get").RGroup().
		Reply(200).
		WJSON(envs)
	mockServer.AddRouter("/nodes/nodeid/update", "post").RGroup().
		Reply(200).
		RQuery("version=100").
		RJSON(requestBody)

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

	err = craneDockerClient.UpdateNode(node, opts)
	assert.Nil(t, err)
	err = craneDockerClient.UpdateNode(node, optsJSONError)
	assert.NotNil(t, err)
}

func TestUpdateNodeEndpointUpdate(t *testing.T) {
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
	node := swarm.Node{
		ID:   "nodeid",
		Spec: swarm.NodeSpec{},
		Meta: swarm.Meta{
			Version: swarm.Version{
				Index: uint64(100),
			},
		},
	}
	opts := model.UpdateOptions{
		Method:  flagEndpointUpdate,
		Options: []byte(`"wrong ip"`),
	}

	mockServer.AddRouter("/_ping", "get").RGroup().
		Reply(200)
	mockServer.AddRouter("/version", "get").RGroup().
		Reply(200).
		WJSON(envs)
	mockServer.AddRouter("/nodes/nodeid/update", "post").RGroup().
		Reply(200).
		RQuery("version=100")

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

	err = craneDockerClient.UpdateNode(node, opts)
	assert.NotNil(t, err)
}

func TestUpdateNodeDefault(t *testing.T) {
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
	node := swarm.Node{
		Spec: swarm.NodeSpec{},
	}
	opts := model.UpdateOptions{
		Method: "defaultError",
	}

	mockServer.AddRouter("/_ping", "get").RGroup().
		Reply(200)
	mockServer.AddRouter("/version", "get").RGroup().
		Reply(200).
		WJSON(envs)
	mockServer.AddRouter("/nodes/nodeid/update", "post").RGroup().
		Reply(200).
		RQuery("version=100")

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

	err = craneDockerClient.UpdateNode(node, opts)
	assert.NotNil(t, err)
}

// TODO (wtzhou) refactor me by test/testing mock
func TestCreateNodeRoleManager(t *testing.T) {
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

	fakeClusterWithID := func(ID string) func(ctx *gin.Context) {
		fakeCluster := func(ctx *gin.Context) {
			var body swarm.Swarm
			body.JoinTokens.Manager = "FakeManagerToken"
			ctx.JSON(http.StatusOK, body)
		}
		return fakeCluster
	}

	fakeNodeInfo := func(addr string, nodeID string) func(ctx *gin.Context) {
		nodeInfo := func(ctx *gin.Context) {
			body := types.Info{
				Swarm: swarm.Info{
					NodeAddr: addr,
					NodeID:   nodeID,
				},
			}
			ctx.JSON(http.StatusOK, body)
		}
		return nodeInfo
	}

	fakeSwarmNodeInfo := func(addr string, nodeID string) func(ctx *gin.Context) {
		swarmNodeInfo := func(ctx *gin.Context) {
			body := swarm.Node{
				ID: nodeID,
				ManagerStatus: &swarm.ManagerStatus{
					Addr: addr,
				},
			}
			ctx.JSON(http.StatusOK, body)
		}
		return swarmNodeInfo
	}

	fakeNodeUpdate := func(ctx *gin.Context) {
		body := ``
		ctx.JSON(http.StatusOK, body)
	}

	fakeSwarmJoin := func(ctx *gin.Context) {
		body := ``
		ctx.JSON(http.StatusOK, body)
	}

	fakeManagerID := "FakeManagerID"
	fakeJoiningNodeID := "FakeJoiningNodeID"

	joiningNodeRouter := gin.New()
	joiningNodeRouter.POST("/swarm/join", fakeSwarmJoin)
	joiningNodeRouter.GET("/info", fakeNodeInfo("FakeAddr", fakeJoiningNodeID))
	joiningNode := httptest.NewServer(joiningNodeRouter)
	defer joiningNode.Close()

	managerRouter := gin.New()
	managerRouter.GET("/swarm", fakeClusterWithID(fakeManagerID))
	managerRouter.GET("/info", fakeNodeInfo("FakeAddr", fakeManagerID))
	managerRouter.GET("/nodes/"+fakeManagerID, fakeSwarmNodeInfo("FakeAddr", fakeManagerID))
	managerRouter.GET("/nodes/"+fakeJoiningNodeID, fakeSwarmNodeInfo(joiningNode.URL, fakeJoiningNodeID))
	managerRouter.POST("/nodes/"+fakeJoiningNodeID+"/update", fakeNodeUpdate)

	manager := httptest.NewServer(managerRouter)
	defer manager.Close()

	joiningNodeRoleManager := model.JoiningNode{
		Role:     swarm.NodeRoleManager,
		Endpoint: joiningNode.URL,
	}

	httpClient, err := NewHttpClient()
	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: manager.URL,
	}

	err = client.CreateNode(joiningNodeRoleManager)
	assert.Nil(t, err)
}

func TestCreateNodeRoleWorker(t *testing.T) {
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

	fakeClusterWithID := func(ID string) func(ctx *gin.Context) {
		fakeCluster := func(ctx *gin.Context) {
			var body swarm.Swarm
			body.JoinTokens.Worker = "FakeWorkerToken"
			ctx.JSON(http.StatusOK, body)
		}
		return fakeCluster
	}

	fakeNodeInfo := func(addr string, nodeID string) func(ctx *gin.Context) {
		nodeInfo := func(ctx *gin.Context) {
			body := types.Info{
				Swarm: swarm.Info{
					NodeAddr: addr,
					NodeID:   nodeID,
				},
			}
			ctx.JSON(http.StatusOK, body)
		}
		return nodeInfo
	}

	fakeSwarmNodeInfo := func(addr string, nodeID string) func(ctx *gin.Context) {
		swarmNodeInfo := func(ctx *gin.Context) {
			body := swarm.Node{
				ID: nodeID,
				ManagerStatus: &swarm.ManagerStatus{
					Addr: addr,
				},
			}
			ctx.JSON(http.StatusOK, body)
		}
		return swarmNodeInfo
	}

	fakeNodeUpdate := func(ctx *gin.Context) {
		body := ``
		ctx.JSON(http.StatusOK, body)
	}

	fakeSwarmJoin := func(ctx *gin.Context) {
		body := ``
		ctx.JSON(http.StatusOK, body)
	}

	fakeManagerID := "FakeManagerID"
	fakeJoiningNodeID := "FakeJoiningNodeID"

	joiningNodeRouter := gin.New()
	joiningNodeRouter.POST("/swarm/join", fakeSwarmJoin)
	joiningNodeRouter.GET("/info", fakeNodeInfo("FakeAddr", fakeJoiningNodeID))
	joiningNode := httptest.NewServer(joiningNodeRouter)
	defer joiningNode.Close()

	managerRouter := gin.New()
	managerRouter.GET("/swarm", fakeClusterWithID(fakeManagerID))
	managerRouter.GET("/info", fakeNodeInfo("FakeAddr", fakeManagerID))
	managerRouter.GET("/nodes/"+fakeManagerID, fakeSwarmNodeInfo("FakeAddr", fakeManagerID))
	managerRouter.GET("/nodes/"+fakeJoiningNodeID, fakeSwarmNodeInfo(joiningNode.URL, fakeJoiningNodeID))
	managerRouter.POST("/nodes/"+fakeJoiningNodeID+"/update", fakeNodeUpdate)

	manager := httptest.NewServer(managerRouter)
	defer manager.Close()

	joiningNodeRoleWorker := model.JoiningNode{
		Role:     swarm.NodeRoleWorker,
		Endpoint: joiningNode.URL,
	}

	httpClient, err := NewHttpClient()
	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: manager.URL,
	}

	err = client.CreateNode(joiningNodeRoleWorker)
	assert.Nil(t, err)
}
func TestCreateNodeWithInvalidRole(t *testing.T) {
	body := `{"Id":"e90302"}`
	swarmManager := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(body))
	}))
	defer swarmManager.Close()

	joiningNodeWithInvalidRole := model.JoiningNode{
		Role:     "invalidRole",
		Endpoint: "invalid",
	}

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)

	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: swarmManager.URL,
	}

	err = client.CreateNode(joiningNodeWithInvalidRole)
	assert.NotNil(t, err)
}

func TestListNodeError(t *testing.T) {
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

	body := `{"Id":"e90302"}`
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(body))
	}))
	defer server1.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)

	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server1.URL,
	}

	_, err = client.ListNode(types.NodeListOptions{})
	assert.NotNil(t, err)
}

func TestListNode(t *testing.T) {
	body := `
	[
		{
    	    "ID":"1t6jojzasio4veexyubvic4j2",
    	    "Version":{
    	        "Index":26607
    	    },
    	    "CreatedAt":"2016-08-26T08:00:24.466491891Z",
    	    "UpdatedAt":"2016-09-08T05:23:49.697933079Z",
    	    "Spec":{
    	        "Labels":{
    	            "dm.reserved.node.endpoint":"http://192.168.59.103:2376"
    	        },
    	        "Role":"worker",
    	        "Availability":"active"
    	    },
    	    "Description":{
    	        "Hostname":"192.168.59.013",
    	        "Platform":{
    	            "Architecture":"x86_64",
    	            "OS":"linux"
    	        },
    	        "Resources":{
    	            "NanoCPUs":2000000000,
    	            "MemoryBytes":3975561216
    	        },
    	        "Engine":{
    	            "EngineVersion":"1.12.0",
    	            "Plugins":[
    	                {
    	                    "Type":"Network",
    	                    "Name":"bridge"
    	                },
    	                {
    	                    "Type":"Network",
    	                    "Name":"host"
    	                },
    	                {
    	                    "Type":"Network",
    	                    "Name":"null"
    	                },
    	                {
    	                    "Type":"Network",
    	                    "Name":"overlay"
    	                },
    	                {
    	                    "Type":"Volume",
    	                    "Name":"local"
    	                }
    	            ]
    	        }
    	    },
    	    "Status":{
    	        "State":"down"
    	    }
    	},
    	{
    	    "ID":"dbspw1g0sjee8ja1khx2w0xtt",
    	    "Version":{
    	        "Index":26603
    	    },
    	    "CreatedAt":"2016-08-26T07:59:50.685235915Z",
    	    "UpdatedAt":"2016-09-08T05:23:36.061728082Z",
    	    "Spec":{
    	        "Labels":{
    	            "dm.reserved.node.endpoint":"192.168.59.104:2376"
    	        },
    	        "Role":"manager",
    	        "Availability":"active"
    	    },
    	    "Description":{
    	        "Hostname":"localhost",
    	        "Platform":{
    	            "Architecture":"x86_64",
    	            "OS":"linux"
    	        },
    	        "Resources":{
    	            "NanoCPUs":2000000000,
    	            "MemoryBytes":3975561216
    	        },
    	        "Engine":{
    	            "EngineVersion":"1.12.0",
    	            "Plugins":[
    	                {
    	                    "Type":"Network",
    	                    "Name":"bridge"
    	                },
    	                {
    	                    "Type":"Network",
    	                    "Name":"host"
    	                },
    	                {
    	                    "Type":"Network",
    	                    "Name":"null"
    	                },
    	                {
    	                    "Type":"Network",
    	                    "Name":"overlay"
    	                },
    	                {
    	                    "Type":"Volume",
    	                    "Name":"local"
    	                }
    	            ]
    	        }
    	    },
    	    "Status":{
    	        "State":"ready"
    	    },
    	    "ManagerStatus":{
    	        "Leader":true,
    	        "Reachability":"reachable",
    	        "Addr":"192.168.59.104:2377"
    	    }
    	}
	]
	`
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(body))
	}))
	defer server1.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)

	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server1.URL,
	}

	nodes, err := client.ListNode(types.NodeListOptions{})
	assert.Nil(t, err)
	assert.Equal(t, len(nodes), 2)
	assert.Equal(t, nodes[0].ID, "1t6jojzasio4veexyubvic4j2")
	assert.Equal(t, nodes[1].ID, "dbspw1g0sjee8ja1khx2w0xtt")
}

func TestRemoveNode(t *testing.T) {
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/nodes/test" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success"))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("failed"))
		}
		return
	}))
	defer server1.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)

	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server1.URL,
	}

	err = client.RemoveNode("tessst")
	assert.NotNil(t, err)

	err = client.RemoveNode("test")
	assert.Nil(t, err)
}

func TestNodeRole(t *testing.T) {
	badRole, err := json.Marshal("test")
	assert.Nil(t, err)
	_, err = nodeRole(badRole)
	assert.NotNil(t, err)

	worker, err := json.Marshal("worker")
	assert.Nil(t, err)
	role, err := nodeRole(worker)
	assert.Nil(t, err)
	assert.Equal(t, role, swarm.NodeRoleWorker)

	manager, err := json.Marshal("manager")
	assert.Nil(t, err)
	role, err = nodeRole(manager)
	assert.Nil(t, err)
	assert.Equal(t, role, swarm.NodeRoleManager)
}

func TestNodeAvailability(t *testing.T) {
	badAvailability, err := json.Marshal("test")
	assert.Nil(t, err)
	_, err = nodeAvailability(badAvailability)
	assert.NotNil(t, err)

	active, err := json.Marshal("active")
	assert.Nil(t, err)
	availability, err := nodeAvailability(active)
	assert.Nil(t, err)
	assert.Equal(t, swarm.NodeAvailabilityActive, availability)

	pause, err := json.Marshal("pause")
	assert.Nil(t, err)
	availability, err = nodeAvailability(pause)
	assert.Nil(t, err)
	assert.Equal(t, swarm.NodeAvailabilityPause, availability)

	drain, err := json.Marshal("drain")
	assert.Nil(t, err)
	availability, err = nodeAvailability(drain)
	assert.Nil(t, err)
	assert.Equal(t, swarm.NodeAvailabilityDrain, availability)
}

func TestGetDaemonUrlByIdErrorKey(t *testing.T) {
	body := `
	{
	    "ID":"1t6jojzasio4veexyubvic4j2",
	    "CreatedAt":"2016-08-26T08:00:24.466491891Z",
	    "UpdatedAt":"2016-09-08T05:23:49.697933079Z",
	    "Spec":{
	        "Labels":{
	            "dm.reserved.node.endpoint":"http://192.168.59.103:2376"
	        },
	        "Role":"worker",
	        "Availability":"active"
	    },
	    "Description":{
	        "Hostname":"192.168.59.013",
	        "Platform":{
	            "Architecture":"x86_64",
	            "OS":"linux"
	        },
	        "Resources":{
	            "NanoCPUs":2000000000,
	            "MemoryBytes":3975561216
	        },
	        "Engine":{
	            "EngineVersion":"1.12.0",
	            "Plugins":[
	                {
	                    "Type":"Network",
	                    "Name":"bridge"
	                }
	            ]
	        }
	    }
	}
	`
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(body))
	}))
	defer server1.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)

	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server1.URL,
	}

	_, err = client.GetDaemonUrlById("test")
	assert.NotNil(t, err)
}

func TestGetDaemonUrlById(t *testing.T) {
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

	body := `
	{
	    "ID":"1t6jojzasio4veexyubvic4j2",
	    "CreatedAt":"2016-08-26T08:00:24.466491891Z",
	    "UpdatedAt":"2016-09-08T05:23:49.697933079Z",
	    "Spec":{
	        "Labels":{
	            "crane.reserved.node.endpoint":"http://192.168.59.103:2376"
	        },
	        "Role":"worker",
	        "Availability":"active"
	    },
	    "Description":{
	        "Hostname":"192.168.59.013",
	        "Platform":{
	            "Architecture":"x86_64",
	            "OS":"linux"
	        },
	        "Resources":{
	            "NanoCPUs":2000000000,
	            "MemoryBytes":3975561216
	        },
	        "Engine":{
	            "EngineVersion":"1.12.0",
	            "Plugins":[
	                {
	                    "Type":"Network",
	                    "Name":"bridge"
	                }
	            ]
	        }
	    }
	}
	`
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(body))
	}))
	defer server1.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)

	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server1.URL,
	}

	_, err = client.GetDaemonUrlById("test")
	assert.Nil(t, err)
}

func TestGetNodeIdByUrl(t *testing.T) {
	body := `
	{
	    "Swarm":{
	        "NodeID":"dbspw1g0sjee8ja1khx2w0xtt"
	    }
	}
	`
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(body))
	}))
	defer server1.Close()

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)

	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: server1.URL,
	}

	u, err := url.Parse(server1.URL)
	assert.Nil(t, err)

	var returnedNodeId string
	matchedNodeUrlWithSchemeTcp := u
	matchedNodeUrlWithSchemeTcp.Scheme = "tcp"
	returnedNodeId, err = client.getNodeIdByUrl(matchedNodeUrlWithSchemeTcp)
	assert.Nil(t, err)
	assert.Equal(t, returnedNodeId, "dbspw1g0sjee8ja1khx2w0xtt")

	matchedNodeUrlWithSchemeHttp := u
	matchedNodeUrlWithSchemeHttp.Scheme = "http"
	returnedNodeId, err = client.getNodeIdByUrl(matchedNodeUrlWithSchemeHttp)
	assert.Nil(t, err)
	assert.Equal(t, returnedNodeId, "dbspw1g0sjee8ja1khx2w0xtt")

	matchedNodeUrlWithoutScheme := u
	matchedNodeUrlWithoutScheme.Scheme = ""
	returnedNodeId, err = client.getNodeIdByUrl(matchedNodeUrlWithoutScheme)
	assert.NotNil(t, err)
	assert.Equal(t, returnedNodeId, "")

	misMatchedNodeUrl := u
	misMatchedNodeUrl.Host = misMatchedNodeUrl.Host + "mis-match"
	returnedNodeId, err = client.getNodeIdByUrl(misMatchedNodeUrl)
	assert.NotNil(t, err)
	assert.Equal(t, returnedNodeId, "")
}
