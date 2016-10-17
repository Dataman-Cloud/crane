package search

import (
	"os"
	"testing"

	"github.com/Dataman-Cloud/crane/src/utils/config"
	"github.com/docker/engine-api/types/swarm"

	"github.com/Dataman-Cloud/crane/src/dockerclient"
	mock "github.com/Dataman-Cloud/crane/src/testing"
	"github.com/stretchr/testify/assert"
)

func TestNewCraneIndex(t *testing.T) {
	craneIndex := NewCraneIndex(&dockerclient.CraneDockerClient{})
	assert.Equal(t, craneIndex.CraneDockerClient, &dockerclient.CraneDockerClient{}, "should be equal")
}

func TestSearchIndex(t *testing.T) {
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
		ID: "24ifsmvkjbyhk",
		Spec: swarm.NodeSpec{
			Role:         "manager",
			Availability: "active",
			Annotations: swarm.Annotations{
				Name: "my-node",
				Labels: map[string]string{
					"crane.reserved.node.endpoint": mockServer.Scheme + "://" + mockServer.Addr + ":" + mockServer.Port,
				},
			},
		},
	}
	nodes := make([]swarm.Node, 0)
	nodes = append(nodes, node)
	networks := `[
		     {
		             "ID": "9fb1e39c",
		             "Name": "foo",
		             "Type": "bridge",
		             "Endpoints":[{"ID": "c080be979dda", "Name": "lllll2222", "Network": "9fb1e39c"}]
		     }
	]`
	nodeInfo := `{
            "Swarm":{
                "NodeID": "24ifsmvkjbyhk"
            }
        }`
	volumes := `[
		{
			"Name": "tardis",
			"Driver": "local",
			"Mountpoint": "/var/lib/docker/volumes/tardis"
		},
		{
			"Name": "foo",
			"Driver": "bar",
			"Mountpoint": "/var/lib/docker/volumes/bar"
		}
	]`
	volumesBody := `{ "Volumes":` + volumes + `}`
	mockServer.AddRouter("/_ping", "get").RGroup().
		Reply(200)
	mockServer.AddRouter("/version", "get").RGroup().
		Reply(200).
		WJSON(envs)
	mockServer.AddRouter("/nodes", "get").RGroup().
		Reply(200).
		WJSON(nodes)
	mockServer.AddRouter("/nodes/24ifsmvkjbyhk", "get").RGroup().
		Reply(200).
		WJSON(node)
	mockServer.AddRouter("/info", "get").RGroup().
		Reply(200).
		WBodyString(nodeInfo)
	mockServer.AddRouter("/networks", "get").RGroup().
		RQuery("filters={}").
		Reply(200).
		WJSON(networks)
	mockServer.AddRouter("/volumes", "get").RGroup().
		Reply(200).
		WJSON(volumesBody)
	mockServer.AddRouter("/services", "get").RGroup().
		RQuery(`filters={"label":{"com.docker.stack.namespace":true}}`).
		Reply(200).
		WFile("./services.json")
	router1 := mockServer.AddRouter("/tasks", "get")
	router1.RGroup().
		RQuery(`filters={"service":{"9mnpnzenvg8p8tdbtq4wvbkcz":true}}`).
		Reply(200).
		WFile("./tasks.json")
	router1.RGroup().
		Reply(200).
		WFile("./tasks.json")
	mockServer.Register()

	config := &config.Config{
		DockerEntryScheme: mockServer.Scheme,
		SwarmManagerIP:    mockServer.Addr,
		DockerEntryPort:   mockServer.Port,
		DockerTlsVerify:   false,
		DockerApiVersion:  "",
	}
	craneDockerClient, err := dockerclient.NewCraneDockerClient(config)
	if err != nil {
		t.Error("fails to create CraneDockerClient:", err)
	}
	craneIndexer := &CraneIndexer{
		CraneDockerClient: craneDockerClient,
	}
	documentStorage := &DocumentStorage{
		Store: map[string]Document{},
	}
	craneIndexer.Index(documentStorage)
}
