package dockerclient

import (
	"bufio"
	"io"

	log "github.com/Sirupsen/logrus"
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

func (client *RolexDockerClient) DiffContainer(containerID string) ([]goclient.Change, error) {
	return client.DockerClient("placeholder").ContainerChanges(containerID)
}

func (client *RolexDockerClient) RenameContainer(opts goclient.RenameContainerOptions) error {
	return client.DockerClient("placeholder").RenameContainer(opts)
}

func (client *RolexDockerClient) StopContainer(containerId string, timeout uint) error {
	return client.DockerClient("placeholder").StopContainer(containerId, timeout)
}

func (client *RolexDockerClient) StartContainer(containerID string, hostconfig *goclient.HostConfig) error {
	return client.DockerClient("placeholder").StartContainer(containerID, hostconfig)
}

func (client *RolexDockerClient) RestartContainer(containerId string, timeout uint) error {
	return client.DockerClient("placeholder").RestartContainer(containerId, timeout)
}

func (client *RolexDockerClient) PauseContainer(containerID string) error {
	return client.DockerClient("placeholder").PauseContainer(containerID)
}

func (client *RolexDockerClient) UnpauseContainer(containerID string) error {
	return client.DockerClient("placeholder").UnpauseContainer(containerID)
}

func (client *RolexDockerClient) ResizeContainerTTY(containerID string, height, width int) error {
	return client.DockerClient("placeholder").ResizeContainerTTY(containerID, height, width)
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
