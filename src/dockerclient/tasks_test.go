package dockerclient

import (
	"testing"
	"time"

	"github.com/docker/engine-api/types/swarm"
	"github.com/stretchr/testify/assert"

	mock "github.com/Dataman-Cloud/crane/src/testing"
	"github.com/Dataman-Cloud/crane/src/utils/config"
)

func TestLen(t *testing.T) {
	tasks := Tasks{
		swarm.Task{
			ID: "0",
		},
		swarm.Task{
			ID: "1",
		},
	}
	assert.Equal(t, tasks.Len(), 2)
}

func TestSwap(t *testing.T) {
	tasks := Tasks{
		swarm.Task{
			ID: "0",
		},
		swarm.Task{
			ID: "1",
		},
	}
	tasks.Swap(0, 1)
	assert.Equal(t, tasks[0].ID, "1")
	assert.Equal(t, tasks[1].ID, "0")
}

func TestLess(t *testing.T) {
	t0 := time.Now()
	t1 := t0.AddDate(1, 1, 1)
	tasks := Tasks{
		swarm.Task{
			Meta: swarm.Meta{
				CreatedAt: t0,
			},
		},
		swarm.Task{
			Meta: swarm.Meta{
				CreatedAt: t1,
			},
		},
	}
	result := tasks.Less(0, 1)

	assert.True(t, result)
}

func TestInspectTask(t *testing.T) {
	mockServer := mock.NewServer()
	defer mockServer.Close()

	task := swarm.Task{
		ID: "faketaskid",
	}
	taskWithWrongJSON := `
	{
		"ID":         "wrongtaskid",
		"WRONG_JSON": "cover err: json Unmarshal",
	}
	`
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

	mockServer.AddRouter("/_ping", "get").RGroup().
		Reply(200)
	mockServer.AddRouter("/version", "get").RGroup().
		Reply(200).
		WJSON(envs)
	mockServer.AddRouter("/tasks/"+task.ID, "get").RGroup().
		Reply(200).
		WJSON(task)
	mockServer.AddRouter("/tasks/wrongtaskid", "get").RGroup().
		Reply(400).
		WJSON(taskWithWrongJSON)

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

	returned, err := craneDockerClient.InspectTask(task.ID)
	assert.Nil(t, err)
	assert.Equal(t, returned.ID, task.ID)

	returned, err = craneDockerClient.InspectTask("wrongtaskid")
	assert.NotNil(t, err)
}
