package model

import (
	"time"

	goclient "github.com/fsouza/go-dockerclient"
)

type ContainerStat struct {
	NodeId      string
	ServiceId   string
	ServiceName string
	TaskId      string
	TaskName    string
	ContainerId string
	Stat        *goclient.Stats
}

type ContainerStatOptions struct {
	ID                  string
	Stats               chan *goclient.Stats
	Stream              bool
	Timeout             time.Duration
	Done                chan bool
	InactivityTimeout   time.Duration
	RolexContainerStats chan *ContainerStat
}
