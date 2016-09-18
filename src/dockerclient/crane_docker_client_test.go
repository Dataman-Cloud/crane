package dockerclient

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Dataman-Cloud/crane/src/utils/config"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	"github.com/stretchr/testify/assert"
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
	client, err := NewGoDockerClientTls("x.x.x.x", "1.23", &config.Config{})
	assert.NotNil(t, err)

	assert.Nil(t, client)
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

	u, err := url.Parse(server1.URL)
	assert.Nil(t, err)

	err = client.VerifyNodeEndpoint("dbspw1g0sjee8ja1khx2w0xtt", u)
	assert.Nil(t, err)
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
