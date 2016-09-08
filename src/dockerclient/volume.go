package dockerclient

import (
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	"golang.org/x/net/context"
)

func (client *CraneDockerClient) InspectVolume(ctx context.Context, name string) (*docker.Volume, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.InspectVolume(name)
}

func (client *CraneDockerClient) ListVolumes(ctx context.Context, opts docker.ListVolumesOptions) ([]docker.Volume, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.ListVolumes(opts)
}

func (client *CraneDockerClient) CreateVolume(ctx context.Context, opts docker.CreateVolumeOptions) (*docker.Volume, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}

	if opts.Name == "" || !isValidName.MatchString(opts.Name) {
		return nil, cranerror.NewError(CodeInvalidVolumeName, "invalid name, only [a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9] are allowed")
	}

	return swarmNode.CreateVolume(opts)
}

func (client *CraneDockerClient) RemoveVolume(ctx context.Context, name string) error {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return err
	}
	return swarmNode.RemoveVolume(name)
}
