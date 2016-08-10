package model

import (
	"time"

	docker "github.com/fsouza/go-dockerclient"
)

type ContainerStat struct {
	NodeId      string
	ServiceId   string
	ServiceName string
	TaskId      string
	TaskName    string
	ContainerId string
	Stat        *docker.Stats
}

type ContainerStatOptions struct {
	ID                  string
	Stats               chan *docker.Stats
	Stream              bool
	Timeout             time.Duration
	Done                chan bool
	InactivityTimeout   time.Duration
	RolexContainerStats chan *ContainerStat
}
