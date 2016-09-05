package model

import (
	"time"

	docker "github.com/Dataman-Cloud/go-dockerclient"
)

type RolexContainerStat struct {
	NodeId      string
	ServiceId   string
	ServiceName string
	TaskId      string
	TaskName    string
	ContainerId string
	Stat        *docker.Stats
	ReceiveRate uint64
	SendRate    uint64
}

type ContainerStatOptions struct {
	ID                  string
	Stats               chan *docker.Stats
	Stream              bool
	Timeout             time.Duration
	Done                chan bool
	InactivityTimeout   time.Duration
	RolexContainerStats chan *RolexContainerStat
}
