package dockerclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
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

func TestListNodeError(t *testing.T) {
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
