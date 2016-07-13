package model

import (
	goclient "github.com/fsouza/go-dockerclient"
)

type ConnectNetwork struct {
	Method         string
	NetworkOptions goclient.NetworkConnectionOptions
}
