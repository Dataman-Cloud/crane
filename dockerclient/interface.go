package dockerclient

// NodeListOptions holds parameters to list  nodes with.
type NodeListOptions struct {
}

type DockerClient interface {
	Ping() error

	// NodeList(opts NodeListOptions) ([]swarm.Node, error)
}
