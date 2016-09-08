package dockerclient

import (
	"encoding/json"
	"net/url"
	"sort"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
)

type Tasks []swarm.Task

func (t Tasks) Len() int {
	return len(t)
}

func (t Tasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Tasks) Less(i, j int) bool {
	return t[j].CreatedAt.Unix() < t[i].CreatedAt.Unix()
}

// TaskList returns the list of tasks.
func (client *CraneDockerClient) ListTasks(options types.TaskListOptions) (Tasks, error) {
	query := url.Values{}

	if options.Filter.Len() > 0 {
		filterJSON, err := filters.ToParam(options.Filter)
		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}

	var tasks Tasks
	content, err := client.sharedHttpClient.GET(nil, client.swarmManagerHttpEndpoint+"/tasks", query, nil)
	if err != nil {
		return tasks, err
	}

	if err := json.Unmarshal(content, &tasks); err != nil {
		return tasks, err
	}

	sort.Sort(tasks)

	return tasks, nil
}

// TaskInspect returns the list of tasks.
func (client *CraneDockerClient) InspectTask(taskID string) (*swarm.Task, error) {
	task := &swarm.Task{}

	content, err := client.sharedHttpClient.GET(nil, client.swarmManagerHttpEndpoint+"/tasks/"+taskID, nil, nil)
	if err != nil {
		return task, err
	}

	if err := json.Unmarshal(content, task); err != nil {
		return task, err
	}

	return task, nil
}
