package model

import (
	goclient "github.com/fsouza/go-dockerclient"
)

type Stats struct {
	NodeId      string
	ServiceId   string
	ServiceName string
	TaskId      string
	TaskName    string
	ContainerId string
	Stat        *goclient.Stats
}
