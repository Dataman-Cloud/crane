package dockerclient

import (
	"bufio"
	"encoding/json"
	"io"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	goclient "github.com/fsouza/go-dockerclient"
)

func (client *RolexDockerClient) ListContainers(opts goclient.ListContainersOptions) ([]goclient.APIContainers, error) {
	return client.DockerClient("placeholder").ListContainers(opts)
}

func (client *RolexDockerClient) InspectContainer(id string) (*goclient.Container, error) {
	return client.DockerClient("placeholder").InspectContainer(id)
}

func (client *RolexDockerClient) RemoveContainer(opts goclient.RemoveContainerOptions) error {
	return client.DockerClient("placeholder").RemoveContainer(opts)
}

func (client *RolexDockerClient) KillContainer(opts goclient.KillContainerOptions) error {
	return client.DockerClient("placeholder").KillContainer(opts)
}

func (client *RolexDockerClient) DiffContainer(containerID string) ([]types.ContainerChange, error) {
	var changes []types.ContainerChange

	content, err := client.HttpGet("/containers/"+containerID+"/changes", nil, nil)
	if err != nil {
		return changes, err
	}

	if err := json.Unmarshal(content, &changes); err != nil {
		return changes, err
	}

	return changes, nil
}

func (client *RolexDockerClient) LogsContainer(nodeId, containerId string, message chan string) {
	outrd, outwr := io.Pipe()
	errrd, errwr := io.Pipe()

	go logReader(outrd, message)
	go logReader(errrd, message)

	opts := goclient.LogsOptions{
		Container:    containerId,
		OutputStream: outwr,
		ErrorStream:  errwr,
		Stdout:       true,
		Stderr:       true,
		Follow:       true,
		Tail:         "0",
	}
	err := client.DockerClient(nodeId).Logs(opts)
	log.Infof("read container log error: %v", err)
}

func logReader(input *io.PipeReader, message chan string) {
	buf := bufio.NewReader(input)

	defer func() {
		if err := recover(); err != nil {
			log.Errorf("send container log to channel error: %v", err)
		}
		input.Close()
		return
	}()

	for {
		select {
		default:
			if line, err := buf.ReadBytes('\n'); err != nil {
				log.Errorf("container log read buffer error: %v", err)
				return
			} else {
				message <- string(line)
			}
		}
	}
}

func (client *RolexDockerClient) StatsContainer(nodeId, containerId string, stats chan *goclient.Stats, done chan bool) {
	opts := goclient.StatsOptions{
		ID:     containerId,
		Stats:  stats,
		Stream: true,
		Done:   done,
	}
	err := client.DockerClient(nodeId).Stats(opts)
	log.Infof("stats of container error: %v", err)
}
