package model

import (
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
