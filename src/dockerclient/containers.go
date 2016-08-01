package dockerclient

import (
	"bufio"
	"io"

	"github.com/Dataman-Cloud/rolex/src/dockerclient/model"

	log "github.com/Sirupsen/logrus"
	goclient "github.com/fsouza/go-dockerclient"
	"golang.org/x/net/context"
)

func (client *RolexDockerClient) ListContainers(ctx context.Context, opts goclient.ListContainersOptions) ([]goclient.APIContainers, error) {
	return client.DockerClient(ctx).ListContainers(opts)
}

func (client *RolexDockerClient) InspectContainer(ctx context.Context, id string) (*goclient.Container, error) {
	return client.DockerClient(ctx).InspectContainer(id)
}

func (client *RolexDockerClient) RemoveContainer(ctx context.Context, opts goclient.RemoveContainerOptions) error {
	return client.DockerClient(ctx).RemoveContainer(opts)
}

func (client *RolexDockerClient) KillContainer(ctx context.Context, opts goclient.KillContainerOptions) error {
	return client.DockerClient(ctx).KillContainer(opts)
}

func (client *RolexDockerClient) RenameContainer(ctx context.Context, opts goclient.RenameContainerOptions) error {
	return client.DockerClient(ctx).RenameContainer(opts)
}

func (client *RolexDockerClient) DiffContainer(ctx context.Context, containerID string) ([]goclient.Change, error) {
	return client.DockerClient(ctx).ContainerChanges(containerID)
}

func (client *RolexDockerClient) StopContainer(ctx context.Context, containerId string, timeout uint) error {
	return client.DockerClient(ctx).StopContainer(containerId, timeout)
}

func (client *RolexDockerClient) StartContainer(ctx context.Context, containerID string, hostconfig *goclient.HostConfig) error {
	return client.DockerClient(ctx).StartContainer(containerID, hostconfig)
}

func (client *RolexDockerClient) RestartContainer(ctx context.Context, containerId string, timeout uint) error {
	return client.DockerClient(ctx).RestartContainer(containerId, timeout)
}

func (client *RolexDockerClient) PauseContainer(ctx context.Context, containerID string) error {
	return client.DockerClient(ctx).PauseContainer(containerID)
}

func (client *RolexDockerClient) UnpauseContainer(ctx context.Context, containerID string) error {
	return client.DockerClient(ctx).UnpauseContainer(containerID)
}

func (client *RolexDockerClient) ResizeContainerTTY(ctx context.Context, containerID string, height, width int) error {
	return client.DockerClient(ctx).ResizeContainerTTY(containerID, height, width)
}

func (client *RolexDockerClient) LogsContainer(ctx context.Context, containerId string, message chan string) {
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
	err := client.DockerClient(ctx).Logs(opts)
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

func (client *RolexDockerClient) StatsContainer(ctx context.Context, containerId string, stats chan *model.ContainerStat) {
	stat := make(chan *goclient.Stats)
	sd := make(chan bool)
	opts := goclient.StatsOptions{
		ID:     containerId,
		Stats:  stat,
		Stream: true,
		Done:   sd,
	}

	container, err := client.DockerClient(ctx).InspectContainer(containerId)
	if err != nil {
		log.Errorf("stats container get container by containerId error: %v", err)
		return
	}
	go func(s chan *goclient.Stats, msg chan *model.ContainerStat, sdone chan bool) {
		defer func() {
			recover()
			sdone <- true
		}()

		for {
			select {
			case data := <-s:
				msg <- &model.ContainerStat{
					Stat:        data,
					NodeId:      container.Config.Labels["com.docker.swarm.node.id"],
					ServiceId:   container.Config.Labels["com.docker.swarm.service.id"],
					ServiceName: container.Config.Labels["com.docker.swarm.service.name"],
					TaskId:      container.Config.Labels["com.docker.swarm.task.id"],
					TaskName:    container.Config.Labels["com.docker.swarm.task.name"],
					ContainerId: container.ID,
				}
			}
		}
	}(stat, stats, sd)

	err = client.DockerClient(ctx).Stats(opts)
	log.Infof("stats of container error: %v", err)
}
