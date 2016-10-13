package dockerclient

import (
	"testing"
	"time"

	"github.com/Dataman-Cloud/crane/src/dockerclient/model"
	mock "github.com/Dataman-Cloud/crane/src/testing"
	"github.com/Dataman-Cloud/crane/src/utils/config"

	"github.com/docker/engine-api/types/swarm"
	"github.com/stretchr/testify/assert"
)

func TestValidateResources(t *testing.T) {
	err := validateResources(&swarm.Resources{})
	assert.Nil(t, err)

	err = validateResources(&swarm.Resources{NanoCPUs: 1e5})
	assert.NotNil(t, err)

	err = validateResources(&swarm.Resources{MemoryBytes: 3 * 1024 * 1024})
	assert.NotNil(t, err)

	err = validateResources(&swarm.Resources{NanoCPUs: 1e6, MemoryBytes: 4 * 1024 * 1024})
	assert.Nil(t, err)
}

func TestValidateResourceRequirements(t *testing.T) {
	err := validateResourceRequirements(&swarm.ResourceRequirements{})
	assert.Nil(t, err)

	err = validateResourceRequirements(&swarm.ResourceRequirements{
		Limits: &swarm.Resources{
			NanoCPUs: 1e5,
		},
	})
	assert.NotNil(t, err)

	err = validateResourceRequirements(&swarm.ResourceRequirements{
		Limits: &swarm.Resources{
			MemoryBytes: 3 * 1024 * 1024,
		},
	})
	assert.NotNil(t, err)

	err = validateResourceRequirements(&swarm.ResourceRequirements{
		Reservations: &swarm.Resources{
			NanoCPUs: 1e5,
		},
	})
	assert.NotNil(t, err)

	err = validateResourceRequirements(&swarm.ResourceRequirements{
		Reservations: &swarm.Resources{
			MemoryBytes: 3 * 1024 * 1024,
		},
	})
	assert.NotNil(t, err)
}

func TestValidateRestartPolicy(t *testing.T) {
	s := time.Duration(time.Second * 10)
	err := validateRestartPolicy(&swarm.RestartPolicy{
		Delay:  &s,
		Window: &s,
	})
	assert.Nil(t, err)

	s = time.Duration(-time.Second * 10)
	err = validateRestartPolicy(&swarm.RestartPolicy{
		Delay: &s,
	})
	assert.NotNil(t, err)

	s = time.Duration(-time.Second * 10)
	err = validateRestartPolicy(&swarm.RestartPolicy{
		Window: &s,
	})
	assert.NotNil(t, err)
}

func TestValidatePlacement(t *testing.T) {
	err := validatePlacement(&swarm.Placement{
		Constraints: []string{"node==1"},
	})
	assert.Nil(t, err)

	err = validatePlacement(&swarm.Placement{
		Constraints: []string{"node=1"},
	})
	assert.NotNil(t, err)
}

func TestValidateUpdate(t *testing.T) {
	s := time.Duration(time.Second * 10)
	err := validateUpdate(&swarm.UpdateConfig{
		Delay: s,
	})
	assert.Nil(t, err)

	s = time.Duration(-time.Second * 10)
	err = validateUpdate(&swarm.UpdateConfig{
		Delay: s,
	})
	assert.NotNil(t, err)
}

func TestValidateTask(t *testing.T) {
	err := validateTask(swarm.TaskSpec{})
	assert.Nil(t, err)

	err = validateTask(swarm.TaskSpec{
		Resources: &swarm.ResourceRequirements{
			Limits: &swarm.Resources{
				NanoCPUs: 1e5,
			},
		},
	})
	assert.NotNil(t, err)

	s := time.Duration(-time.Second * 10)
	err = validateTask(swarm.TaskSpec{
		RestartPolicy: &swarm.RestartPolicy{
			Delay: &s,
		},
	})
	assert.NotNil(t, err)

	err = validateTask(swarm.TaskSpec{
		Placement: &swarm.Placement{
			Constraints: []string{"node=1"},
		},
	})
	assert.NotNil(t, err)
}

func TestValidateEndpointSpec(t *testing.T) {
	err := validateEndpointSpec(&swarm.EndpointSpec{})
	assert.Nil(t, err)

	err = validateEndpointSpec(&swarm.EndpointSpec{
		Mode: swarm.ResolutionModeDNSRR,
		Ports: []swarm.PortConfig{swarm.PortConfig{
			Name:          "test",
			Protocol:      swarm.PortConfigProtocolTCP,
			TargetPort:    uint32(8080),
			PublishedPort: uint32(8080),
		}},
	})
	assert.NotNil(t, err)

	err = validateEndpointSpec(&swarm.EndpointSpec{
		Mode: swarm.ResolutionModeDNSRR,
		Ports: []swarm.PortConfig{
			swarm.PortConfig{
				Name:          "test",
				Protocol:      swarm.PortConfigProtocolTCP,
				TargetPort:    uint32(8080),
				PublishedPort: uint32(8080),
			},
			swarm.PortConfig{
				Name:          "test1",
				Protocol:      swarm.PortConfigProtocolTCP,
				TargetPort:    uint32(8080),
				PublishedPort: uint32(8080),
			},
		},
	})
	assert.NotNil(t, err)
}

func TestValidateCraneServiceSpec(t *testing.T) {
	err := ValidateCraneServiceSpec(&model.CraneServiceSpec{
		Name: "test",
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: swarm.ContainerSpec{
				Image: "testimage:latest",
			},
		},
		UpdateConfig: &swarm.UpdateConfig{},
		EndpointSpec: &swarm.EndpointSpec{},
	})
	assert.Nil(t, err)

	err = ValidateCraneServiceSpec(&model.CraneServiceSpec{
		Name:         "t==-//?",
		TaskTemplate: swarm.TaskSpec{},
		UpdateConfig: &swarm.UpdateConfig{},
		EndpointSpec: &swarm.EndpointSpec{},
	})
	assert.NotNil(t, err)

	s := time.Duration(-time.Second * 10)
	err = ValidateCraneServiceSpec(&model.CraneServiceSpec{
		Name: "test",
		TaskTemplate: swarm.TaskSpec{
			RestartPolicy: &swarm.RestartPolicy{
				Delay: &s,
			},
		},
		UpdateConfig: &swarm.UpdateConfig{},
		EndpointSpec: &swarm.EndpointSpec{},
	})
	assert.NotNil(t, err)

	err = ValidateCraneServiceSpec(&model.CraneServiceSpec{
		Name: "test",
		UpdateConfig: &swarm.UpdateConfig{
			Delay: s,
		},
		EndpointSpec: &swarm.EndpointSpec{},
	})
	assert.NotNil(t, err)

	err = ValidateCraneServiceSpec(&model.CraneServiceSpec{
		Name: "test",
		EndpointSpec: &swarm.EndpointSpec{
			Mode: swarm.ResolutionModeDNSRR,
			Ports: []swarm.PortConfig{swarm.PortConfig{
				Name:          "test",
				Protocol:      swarm.PortConfigProtocolTCP,
				TargetPort:    uint32(8080),
				PublishedPort: uint32(8080),
			}},
		},
	})
	assert.NotNil(t, err)

	err = ValidateCraneServiceSpec(&model.CraneServiceSpec{
		Name: "test",
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: swarm.ContainerSpec{
				Image: "sdfg-=-asd",
			},
		},
	})
	assert.NotNil(t, err)
}

func TestValidateName(t *testing.T) {
	err := validateName("testname")
	assert.Nil(t, err)

	err = validateName("")
	assert.NotNil(t, err)

	err = validateName("-=-sdfg")
	assert.NotNil(t, err)
}

func TestValidateImageName(t *testing.T) {
	err := validateImageName("testimage:latest")
	assert.Nil(t, err)

	err = validateImageName("-=-sdfg")
	assert.NotNil(t, err)
}

func TestCheckPortConflicts(t *testing.T) {
	err := checkPortConflicts(map[string]bool{"test": true}, "test", []swarm.Service{
		swarm.Service{
			ID: "test1",
			Spec: swarm.ServiceSpec{
				EndpointSpec: &swarm.EndpointSpec{
					Ports: []swarm.PortConfig{swarm.PortConfig{
						Protocol:      swarm.PortConfigProtocolTCP,
						TargetPort:    uint32(8080),
						PublishedPort: uint32(8080),
					}},
				},
			},
		},
	})
	assert.Nil(t, err)

	err = checkPortConflicts(map[string]bool{"8080/tcp": true}, "", []swarm.Service{
		swarm.Service{
			ID: "test1",
			Spec: swarm.ServiceSpec{
				EndpointSpec: &swarm.EndpointSpec{
					Ports: []swarm.PortConfig{swarm.PortConfig{
						Protocol:      swarm.PortConfigProtocolTCP,
						TargetPort:    uint32(8080),
						PublishedPort: uint32(8080),
					}},
				},
			},
		},
	})
	assert.NotNil(t, err)

	err = checkPortConflicts(map[string]bool{"8080/tcp": true}, "", []swarm.Service{
		swarm.Service{
			ID: "test1",
			Endpoint: swarm.Endpoint{
				Ports: []swarm.PortConfig{swarm.PortConfig{
					Protocol:      swarm.PortConfigProtocolTCP,
					TargetPort:    uint32(8080),
					PublishedPort: uint32(8080),
				}},
			},
		},
	})
	assert.NotNil(t, err)
}

func TestPortConflictToString(t *testing.T) {
	str := PortConflictToString(swarm.PortConfig{
		Protocol:      swarm.PortConfigProtocolTCP,
		PublishedPort: uint32(8080),
	})
	assert.Equal(t, str, "8080/tcp")
}

func TestCheckServicePortConflictsNoServer(t *testing.T) {
	nilEndpointSpec := &model.CraneServiceSpec{
		EndpointSpec: nil,
	}
	serviceId := "serviceId"

	cliNoServer := &CraneDockerClient{}
	returned := cliNoServer.CheckServicePortConflicts(nilEndpointSpec, serviceId)
	assert.Nil(t, returned)

	endPointSpecNoPublishedPort := &model.CraneServiceSpec{
		EndpointSpec: &swarm.EndpointSpec{
			Ports: []swarm.PortConfig{
				swarm.PortConfig{
					PublishedPort: 0,
				},
			},
		},
	}
	returned = cliNoServer.CheckServicePortConflicts(endPointSpecNoPublishedPort, serviceId)
	assert.Nil(t, returned)
}

func TestCheckServicePortConflicts(t *testing.T) {
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

	endPointSpec80 := &swarm.EndpointSpec{
		Ports: []swarm.PortConfig{
			swarm.PortConfig{
				PublishedPort: 80,
			},
		},
	}
	serviceSpec := &model.CraneServiceSpec{
		EndpointSpec: endPointSpec80,
	}
	servicesExisted := []swarm.Service{
		swarm.Service{
			ID: "serviceid",
			Spec: swarm.ServiceSpec{
				EndpointSpec: endPointSpec80,
			},
		},
	}

	mockServer.AddRouter("/_ping", "get").RGroup().
		Reply(200)
	mockServer.AddRouter("/version", "get").RGroup().
		Reply(200).
		WJSON(envs)
	mockServer.AddRouter("/services", "get").RGroup().
		Reply(200).
		WJSON(servicesExisted)

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

	requestedService := serviceSpec
	requestedService.EndpointSpec.Ports[0].PublishedPort = 81
	returned := craneDockerClient.CheckServicePortConflicts(requestedService, "serviceId")
	assert.Nil(t, returned)
}
