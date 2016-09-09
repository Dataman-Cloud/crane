package dockerclient

import (
	"testing"
	"time"

	"github.com/Dataman-Cloud/crane/src/dockerclient/model"

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
