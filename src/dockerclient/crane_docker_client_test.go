package dockerclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/Dataman-Cloud/crane/src/model"
	"github.com/Dataman-Cloud/crane/src/utils/config"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	dockertest "github.com/Dataman-Cloud/go-dockerclient/testing"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestNewHttpClient(t *testing.T) {
	client, err := NewHttpClient()
	assert.Nil(t, err)

	assert.Equal(t, client.HttpClient.Timeout, defaultHttpRequestTimeout)
}

func TestNewHttpClientError(t *testing.T) {
	client, err := NewHttpClientTls(&config.Config{})
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestNewGoDockerClientTlsError(t *testing.T) {
	client, err := NewGoDockerClientTls("x.x.x.x", &config.Config{})
	assert.NotNil(t, err)

	assert.Nil(t, client)
}

func TestNewCraneDockerClient(t *testing.T) {
	conf := &config.Config{}
	craneClient, err := NewCraneDockerClient(conf)
	assert.NotNil(t, err)
	assert.Nil(t, craneClient)

	tlsConf := &config.Config{DockerTlsVerify: true}
	tlsClient, err := NewCraneDockerClient(tlsConf)
	assert.NotNil(t, err)
	assert.Nil(t, tlsClient)
}

func TestVerifyNodeEndpoint(t *testing.T) {
	body := `
	{
	    "ID":"PD5B:R7W4:Q64C:X5VW:ET6Q:OFQP:GFYE:5NYG:V3DU:XS5Q:B6UL:RWIV",
	    "Containers":15,
	    "ContainersRunning":5,
	    "ContainersPaused":0,
	    "ContainersStopped":10,
	    "Images":153,
	    "NCPU":2,
	    "Swarm":{
	        "NodeID":"dbspw1g0sjee8ja1khx2w0xtt",
	        "NodeAddr":"192.168.59.104",
	        "LocalNodeState":"active",
	        "ControlAvailable":true,
	        "Error":"",
	        "RemoteManagers":[
	            {
	                "NodeID":"dbspw1g0sjee8ja1khx2w0xtt",
	                "Addr":"192.168.59.104:2377"
	            }
	        ],
	        "Nodes":2,
	        "Managers":1,
	        "Cluster":{
	            "ID":"edn9ll15c24jxwhvqku8x6p25",
	            "Version":{
	                "Index":26609
	            }
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

	err = client.VerifyNodeEndpoint("test", nil)
	assert.NotNil(t, err)

	u, err := url.Parse(server1.URL)
	assert.Nil(t, err)

	err = client.VerifyNodeEndpoint("dbspw1g0sjee8ja1khx2w0xtt", u)
	assert.Nil(t, err)

	u.Host = "errorHost"
	err = client.VerifyNodeEndpoint("dbspw1g0sjee8ja1khx2w0xtt", u)
	assert.NotNil(t, err)
}

func TestVerifyNodeEndpointErrorId(t *testing.T) {
	body := `
	{
	    "ID":"PD5B:R7W4:Q64C:X5VW:ET6Q:OFQP:GFYE:5NYG:V3DU:XS5Q:B6UL:RWIV",
	    "Containers":15,
	    "ContainersRunning":5,
	    "ContainersPaused":0,
	    "ContainersStopped":10,
	    "Images":153,
	    "NCPU":2,
	    "Swarm":{
	        "NodeID":"dbspw1g0sjee8ja1khx2w0xtt",
	        "NodeAddr":"192.168.59.104",
	        "LocalNodeState":"active",
	        "ControlAvailable":true,
	        "Error":"",
	        "RemoteManagers":[
	            {
	                "NodeID":"dbspw1g0sjee8ja1khx2w0xtt",
	                "Addr":"192.168.59.104:2377"
	            }
	        ],
	        "Nodes":2,
	        "Managers":1,
	        "Cluster":{
	            "ID":"edn9ll15c24jxwhvqku8x6p25",
	            "Version":{
	                "Index":26609
	            }
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

	u, err := url.Parse(server1.URL)
	assert.Nil(t, err)

	err = client.VerifyNodeEndpoint("test", u)
	assert.NotNil(t, err)
}

func TestSwarmManager(t *testing.T) {
	fakeSwarmManager := &docker.Client{}
	fakeCraneDockerClient := &CraneDockerClient{
		swarmManager: fakeSwarmManager,
	}
	swarmManager := fakeCraneDockerClient.SwarmManager()
	assert.Equal(t, fakeSwarmManager, swarmManager)
}

func StartTestServer() *dockertest.DockerServer {
	rand.Seed(time.Now().Unix())
	serverPort := rand.Intn(39999-30000) + 30000
	bindAddr := "127.0.0.1:" + strconv.FormatInt(int64(serverPort), 10)
	testServer, _ := dockertest.NewServer(bindAddr, nil, nil)
	time.Sleep(time.Second)
	return testServer
}

func InitTestSwarm(t *testing.T) (*dockertest.DockerServer, *CraneDockerClient, string) {
	testServer := StartTestServer()
	assert.NotNil(t, testServer)

	httpClient, err := NewHttpClient()
	assert.Nil(t, err)
	endpoint := testServer.URL()[0 : len(testServer.URL())-1]
	managerClient, err := docker.NewVersionedClient(endpoint, "")
	client := &CraneDockerClient{
		sharedHttpClient:         httpClient,
		swarmManagerHttpEndpoint: endpoint,
		swarmManager:             managerClient,
	}

	swarmInitInfo := swarm.InitRequest{
		ForceNewCluster: false,
	}
	_, err = client.sharedHttpClient.POST(nil, client.swarmManagerHttpEndpoint+"/swarm/init", nil, swarmInitInfo, nil)
	assert.Nil(t, err)
	nodes, err := client.ListNode(types.NodeListOptions{})
	assert.Equal(t, len(nodes), 1)

	node := nodes[0]
	var nodeUpdate model.UpdateOptions
	updateOptions := `{"Method":"endpoint-update", "Options": "%s"}`
	updateOptions = fmt.Sprintf(updateOptions, endpoint)
	err = json.Unmarshal([]byte(updateOptions), &nodeUpdate)
	assert.Nil(t, err)

	err = client.UpdateNode(node.ID, nodeUpdate)
	assert.Nil(t, err)
	return testServer, client, node.ID
}

func TestSwarmNode(t *testing.T) {
	testServer, craneClient, nodeId := InitTestSwarm(t)
	assert.NotNil(t, nodeId)
	assert.NotNil(t, craneClient)
	defer testServer.Stop()
	backgroundContext := context.Background()
	_, err := craneClient.SwarmNode(backgroundContext)
	assert.NotNil(t, err)

	craneContext := context.WithValue(backgroundContext, "node_id", "errorId")
	dockerClient, err := craneClient.SwarmNode(craneContext)
	assert.NotNil(t, err)

	craneContext = context.WithValue(backgroundContext, "node_id", nodeId)
	dockerClient, err = craneClient.SwarmNode(craneContext)
	assert.Nil(t, err)
	assert.NotNil(t, dockerClient)
}

func TestToCraneError(t *testing.T) {
	noSuchContainerErr := &docker.NoSuchContainer{ID: "test", Err: errors.New("test error")}
	err := ToCraneError(noSuchContainerErr)
	assert.NotNil(t, err)
	craneErr, ok := err.(*cranerror.CraneError)
	assert.True(t, ok)
	assert.Equal(t, CodeContainerInvalid, craneErr.Code)

	noSuchNetworkErr := &docker.NoSuchNetwork{ID: "test"}
	err = ToCraneError(noSuchNetworkErr)
	assert.NotNil(t, err)
	craneErr, ok = err.(*cranerror.CraneError)
	assert.True(t, ok)
	assert.Equal(t, CodeNetworkInvalid, craneErr.Code)

	noSuchContainerOrNetworkErr := &docker.NoSuchNetworkOrContainer{
		NetworkID:   "test",
		ContainerID: "test",
	}
	err = ToCraneError(noSuchContainerOrNetworkErr)
	assert.NotNil(t, err)
	craneErr, ok = err.(*cranerror.CraneError)
	assert.True(t, ok)
	assert.Equal(t, CodeNetworkOrContainerInvalid, craneErr.Code)

	containerAlreadyRunningErr := &docker.ContainerAlreadyRunning{ID: "test"}
	err = ToCraneError(containerAlreadyRunningErr)
	assert.NotNil(t, err)
	craneErr, ok = err.(*cranerror.CraneError)
	assert.True(t, ok)
	assert.Equal(t, CodeContainerAlreadyRunning, craneErr.Code)

	containerNotRunningErr := &docker.ContainerNotRunning{ID: "test"}
	err = ToCraneError(containerNotRunningErr)
	assert.NotNil(t, err)
	craneErr, ok = err.(*cranerror.CraneError)
	assert.True(t, ok)
	assert.Equal(t, CodeContainerNotRunning, craneErr.Code)

	err = ToCraneError(errors.New("test"))
	assert.NotNil(t, err)
}
