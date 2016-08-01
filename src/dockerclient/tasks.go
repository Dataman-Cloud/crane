package dockerclient

import (
	"encoding/json"
	"net/url"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
)

// TaskList returns the list of tasks.
func (client *RolexDockerClient) ListTasks(options types.TaskListOptions) ([]swarm.Task, error) {
	query := url.Values{}

	if options.Filter.Len() > 0 {
		filterJSON, err := filters.ToParam(options.Filter)
		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}

	var tasks []swarm.Task
	content, err := client.HttpGet(client.SwarmHttpEndpoint+"/tasks", query, nil)
	if err != nil {
		return tasks, err
	}

	if err := json.Unmarshal(content, &tasks); err != nil {
		return tasks, err
	}

	return tasks, nil
}

// TaskInspect returns the list of tasks.
func (client *RolexDockerClient) InspectTask(taskID string) (*swarm.Task, error) {
	task := &swarm.Task{}

	content, err := client.HttpGet(client.SwarmHttpEndpoint+"/tasks/"+taskID, nil, nil)
	if err != nil {
		return task, nil
	}

	if err := json.Unmarshal(content, task); err != nil {
		return task, err
	}

	return task, nil
}
