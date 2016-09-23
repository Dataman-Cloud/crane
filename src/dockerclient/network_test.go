package dockerclient

import (
	"encoding/json"
	"net/http"
	"testing"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestNetwork(t *testing.T) {
	testServer, craneClient, nodeId := InitTestSwarm(t)
	assert.NotNil(t, nodeId)
	assert.NotNil(t, craneClient)
	defer testServer.Stop()
	backgroundContext := context.Background()
	craneContext := context.WithValue(backgroundContext, "node_id", "errorid")
	_, err := craneClient.CreateNodeNetwork(craneContext, docker.CreateNetworkOptions{})
	assert.NotNil(t, err)
	_, err = craneClient.ListNodeNetworks(craneContext, docker.NetworkFilterOpts{})
	assert.NotNil(t, err)
	_, err = craneClient.InspectNodeNetwork(craneContext, "test")
	assert.NotNil(t, err)
	err = craneClient.ConnectNodeNetwork(craneContext, "test", docker.NetworkConnectionOptions{})
	assert.NotNil(t, err)
	err = craneClient.DisconnectNodeNetwork(craneContext, "test", docker.NetworkConnectionOptions{})
	assert.NotNil(t, err)

	craneContext = context.WithValue(backgroundContext, "node_id", nodeId)
	testServer.CustomHandler("/networks", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var networks []docker.Network
		network1 := docker.Network{
			ID:     "test1",
			Name:   "network1",
			Labels: map[string]string{labelNamespace: "stack1"},
		}
		networks = append(networks, network1)
		network2 := docker.Network{
			ID:     "test2",
			Name:   "network2",
			Labels: map[string]string{labelNamespace: "stack2"},
		}
		networks = append(networks, network2)
		network3 := docker.Network{
			ID:   "test3",
			Name: "network3",
		}
		networks = append(networks, network3)

		if r.Method == "GET" {
			if r.URL.Path == "/networks" {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(networks)
			} else {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(network1)
			}
		}

		if r.Method == "DELETE" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(nil)
		}

		if r.URL.Path == "/networks/create" && r.Method == "POST" {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(network1)
		}

		if r.URL.Path == "/networks/test/connect" {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(network1)
		}

		if r.URL.Path == "/networks/test/disconnect" {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(network1)
		}
	}))

	network, err := craneClient.CreateNetwork(docker.CreateNetworkOptions{Name: "@@@#####"})
	assert.NotNil(t, err)
	assert.Nil(t, network)

	network, err = craneClient.CreateNetwork(docker.CreateNetworkOptions{Name: "networks1"})
	assert.Nil(t, err)
	assert.NotNil(t, network)

	networks, err := craneClient.ListNetworks(docker.NetworkFilterOpts{})
	assert.Nil(t, err)
	assert.NotNil(t, networks)
	assert.Equal(t, 3, len(networks))

	networkId := networks[0].ID
	network, err = craneClient.InspectNetwork(networkId)
	assert.Nil(t, err)
	assert.NotNil(t, network)

	err = craneClient.RemoveNetwork(networkId)
	assert.Nil(t, err)

	err = craneClient.ConnectNetwork("test", docker.NetworkConnectionOptions{})
	assert.Nil(t, err)

	err = craneClient.DisconnectNetwork("test", docker.NetworkConnectionOptions{})
	assert.Nil(t, err)

	network, err = craneClient.CreateNodeNetwork(craneContext, docker.CreateNetworkOptions{Name: "@@@#####"})
	assert.NotNil(t, err)
	assert.Nil(t, network)

	network, err = craneClient.CreateNodeNetwork(craneContext, docker.CreateNetworkOptions{Name: "networks1"})
	assert.Nil(t, err)
	assert.NotNil(t, network)

	networks, err = craneClient.ListNodeNetworks(craneContext, docker.NetworkFilterOpts{})
	assert.Nil(t, err)
	assert.NotNil(t, networks)
	assert.Equal(t, 3, len(networks))

	networkId = networks[0].ID
	network, err = craneClient.InspectNodeNetwork(craneContext, networkId)
	assert.Nil(t, err)
	assert.NotNil(t, network)

	err = craneClient.ConnectNodeNetwork(craneContext, "test", docker.NetworkConnectionOptions{})
	assert.Nil(t, err)

	err = craneClient.DisconnectNodeNetwork(craneContext, "test", docker.NetworkConnectionOptions{})
	assert.Nil(t, err)

}
