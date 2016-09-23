package dockerclient

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/Dataman-Cloud/crane/src/dockerclient/model"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"github.com/stretchr/testify/assert"
)

func TestConvertNetworks(t *testing.T) {
	networkMap := map[string]bool{"test1": true, "test2": false}
	networks := []string{"test1", "test2"}
	namespace := "stack"
	name := "service"

	nets := convertNetworks(networkMap, networks, namespace, name)
	assert.NotNil(t, nets)
	assert.Equal(t, 2, len(nets))
	for _, net := range nets {
		assert.Equal(t, 1, len(net.Aliases))
		assert.Equal(t, name, net.Aliases[0])
		if strings.Contains(net.Target, "test1") {
			assert.Equal(t, namespace+"_test1", net.Target)
		}
		if strings.Contains(net.Target, "test2") {
			assert.Equal(t, "test2", net.Target)
		}
	}
}

func TestDeployStack(t *testing.T) {
	testServer, craneClient, _ := InitTestSwarm(t)
	assert.NotNil(t, craneClient)
	defer testServer.Stop()

	testServer.CustomHandler("/services", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(nil)
		}

		if r.URL.Path == "/services/create" && r.Method == "POST" {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(types.ServiceCreateResponse{ID: "service1"})
		}
	}))

	craneServiceSpec := model.CraneServiceSpec{
		Name: "service1",
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: swarm.ContainerSpec{
				Image: "test",
			},
		},
		EndpointSpec: &swarm.EndpointSpec{
			Mode: "vip",
			Ports: []swarm.PortConfig{
				swarm.PortConfig{
					Name:          "stack11",
					Protocol:      swarm.PortConfigProtocolTCP,
					TargetPort:    8888,
					PublishedPort: 9999,
				},
			},
		},
	}
	rightbundle := &model.Bundle{
		Namespace: "stack1",
		Stack: model.BundleService{
			Services: map[string]model.CraneServiceSpec{"service1": craneServiceSpec},
		},
	}

	err := craneClient.DeployStack(rightbundle)
	assert.Nil(t, err)
}

func TestStack(t *testing.T) {
	testServer, craneClient, _ := InitTestSwarm(t)
	assert.NotNil(t, craneClient)
	defer testServer.Stop()

	testServer.CustomHandler("/services", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var services []swarm.Service
			service1 := swarm.Service{
				ID: "test1",
				Spec: swarm.ServiceSpec{
					Annotations: swarm.Annotations{
						Name:   "service1",
						Labels: map[string]string{labelNamespace: "stack1"},
					},
				},
			}
			services = append(services, service1)

			service2 := swarm.Service{
				ID: "test1",
				Spec: swarm.ServiceSpec{
					Annotations: swarm.Annotations{
						Name:   "service1",
						Labels: map[string]string{labelNamespace: "stack1"},
					},
				},
			}
			services = append(services, service2)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(services)
		}

		if r.Method == "delete" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(nil)
		}
	}))

	testServer.CustomHandler("/networks", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
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
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(networks)
		}

		if r.Method == "delete" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(nil)
		}
	}))

	testServer.CustomHandler("/tasks", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var tasks []swarm.Task
			task1 := swarm.Task{
				ID:        "test1",
				ServiceID: "service1",
				Status:    swarm.TaskStatus{State: "running"},
			}
			tasks = append(tasks, task1)
			task2 := swarm.Task{
				ID:        "test2",
				ServiceID: "service1",
				Status:    swarm.TaskStatus{State: "running"},
			}
			tasks = append(tasks, task2)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tasks)
		}
	}))

	testServer.CustomHandler("/nodes", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var nodes []swarm.Node
			node1 := swarm.Node{
				ID:     "test1",
				Status: swarm.NodeStatus{State: "ready"},
			}
			nodes = append(nodes, node1)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(nodes)
		}
	}))

	bundle, err := craneClient.InspectStack("stack1")
	assert.Nil(t, err)
	assert.NotNil(t, bundle)

	stacks, err := craneClient.ListStack()
	assert.Nil(t, err)
	assert.NotNil(t, stacks)

	serviceStatus, err := craneClient.ListStackService("stack1", types.ServiceListOptions{})
	assert.Nil(t, err)
	assert.NotNil(t, serviceStatus)

	err = craneClient.RemoveStack("stack1")
	assert.Nil(t, err)
}
