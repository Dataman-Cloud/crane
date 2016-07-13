package dockerclient

import (
	"encoding/json"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
)

// TaskList returns the list of tasks.
func (client *RolexDockerClient) ListTasks(options types.TaskListOptions) ([]swarm.Task, error) {
	var tasks []swarm.Task

	content, err := client.HttpGet("/tasks")
	if err != nil {
		return tasks, err
	}

	if err := json.Unmarshal(content, &tasks); err != nil {
		return tasks, err
	}

	return tasks, nil
}
