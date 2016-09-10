package dockerclient

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
