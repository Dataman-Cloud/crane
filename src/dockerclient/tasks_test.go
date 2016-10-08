package dockerclient

import (
	"testing"

	"github.com/docker/engine-api/types/swarm"
	"github.com/stretchr/testify/assert"
)

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
