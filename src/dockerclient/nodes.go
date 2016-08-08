package dockerclient

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/Dataman-Cloud/rolex/src/model"
	"github.com/Dataman-Cloud/rolex/src/util/config"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"

	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	goclient "github.com/fsouza/go-dockerclient"
	"golang.org/x/net/context"
)

const (
	flagUpdateRole         = "role"
	flagUpdateAvailability = "availability"
	flagLabelAdd           = "label-add"
	flagLabelRemove        = "label-rm"
)

// NodeList returns the list of nodes.
func (client *RolexDockerClient) ListNode(opts types.NodeListOptions) ([]swarm.Node, error) {
	var nodes []swarm.Node

	content, err := client.HttpGet(client.swarmManagerHttpEndpoint+"/nodes", nil, nil)
	if err != nil {
		return nodes, err
	}

	if err := json.Unmarshal(content, &nodes); err != nil {
		return nodes, err
	}

	return nodes, nil
}

// Inspect node returns the single node.
func (client *RolexDockerClient) InspectNode(nodeId string) (swarm.Node, error) {
	var node swarm.Node

	content, err := client.HttpGet(client.swarmManagerHttpEndpoint+"/"+path.Join("nodes", nodeId), nil, nil)
	if err != nil {
		return node, err
	}

	if err := json.Unmarshal(content, &node); err != nil {
		return node, err
	}

	return node, nil
}

// Remove a single node
func (client *RolexDockerClient) RemoveNode(nodeId string) error {
	_, err := client.HttpDelete(client.swarmManagerHttpEndpoint + "/" + path.Join("nodes", nodeId))
	if err != nil {
		return err
	}

	return nil
}

// Update a single node
func (client *RolexDockerClient) UpdateNode(nodeId string, opts model.UpdateOptions) error {
	node, err := client.InspectNode(nodeId)
	if err != nil {
		return err
	}

	spec := &node.Spec
	switch opts.Method {
	case flagUpdateRole:
		role, err := nodeRole(opts.Options)
		if err != nil {
			return err
		}
		spec.Role = role
	case flagUpdateAvailability:
		availability, err := nodeAvailability(opts.Options)
		if err != nil {
			return err
		}
		spec.Availability = availability
	case flagLabelAdd:
		if err := nodeAddLabels(spec, opts.Options); err != nil {
			return err
		}
	case flagLabelRemove:
		if err := nodeRemoveLabels(spec, opts.Options); err != nil {
			return err
		}
	default:
		errMsg := fmt.Sprintf("Invalid update node method %s", opts.Method)
		return rolexerror.NewRolexError(rolexerror.CodeErrorUpdateNodeMethod, errMsg)
	}

	query := url.Values{}
	query.Set("version", strconv.FormatUint(node.Version.Index, 10))
	_, err = client.HttpPost(client.swarmManagerHttpEndpoint+"/"+path.Join("nodes", nodeId, "update"), query, node.Spec, nil)
	if err != nil {
		return err
	}

	return nil
}

func nodeRole(rawMessage []byte) (swarm.NodeRole, error) {
	var err error
	var role swarm.NodeRole
	if err = json.Unmarshal(rawMessage, &role); err != nil {
		return role, err
	}

	if role != swarm.NodeRoleWorker && role != swarm.NodeRoleManager {
		errMsg := fmt.Sprintf("node role only support %s/%s but got %s",
			swarm.NodeRoleWorker, swarm.NodeRoleManager, role)
		err = rolexerror.NewRolexError(rolexerror.CodeErrorNodeRole, errMsg)
	}

	return role, err
}

func nodeAvailability(rawMessage []byte) (swarm.NodeAvailability, error) {
	var err error
	var availability swarm.NodeAvailability
	if err = json.Unmarshal(rawMessage, &availability); err != nil {
		return availability, err
	}

	if availability != swarm.NodeAvailabilityActive && availability != swarm.NodeAvailabilityPause && availability != swarm.NodeAvailabilityDrain {
		errMsg := fmt.Sprintf("node availability only support %s/%s/%s, but got %s",
			swarm.NodeAvailabilityActive, swarm.NodeAvailabilityPause, swarm.NodeAvailabilityDrain, availability)
		err = rolexerror.NewRolexError(rolexerror.CodeErrorNodeAvailability, errMsg)
	}

	return availability, err
}

func nodeAddLabels(spec *swarm.NodeSpec, rawMessage []byte) error {
	var labelsAdd map[string]string
	if err := json.Unmarshal(rawMessage, &labelsAdd); err != nil {
		return err
	}

	if spec.Annotations.Labels == nil {
		spec.Annotations.Labels = make(map[string]string)
	}

	for k, v := range labelsAdd {
		spec.Annotations.Labels[k] = v
	}

	return nil
}

func nodeRemoveLabels(spec *swarm.NodeSpec, rawMessage []byte) error {
	var labelsRemove []string
	if err := json.Unmarshal(rawMessage, &labelsRemove); err != nil {
		return err
	}

	for _, k := range labelsRemove {
		delete(spec.Annotations.Labels, k)
	}

	return nil
}

// docker info
func (client *RolexDockerClient) Info(ctx context.Context) (*goclient.DockerInfo, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.Info()
}

func (client *RolexDockerClient) NodeDaemonEndpoint(nodeId string, protocol string) (string, error) {
	//node, err := client.InspectNode(nodeId)
	//if err != nil {
	//	return "", err
	//}
	//nodeIp := strings.Split(node.ManagerStatus.Addr, ":")[0]
	//TODO Temporary solution
	var nodeIp string
	for key, val := range config.NodeAddrMap {
		if nodeId == key {
			nodeIp = val
			break
		}
	}
	switch strings.ToLower(protocol) {
	case "http":
		return "http://" + nodeIp + ":" + client.config.NodePort, nil
	case "https":
		return "https://" + nodeIp + ":" + client.config.NodePort, nil
	case "tcp":
		return "tcp://" + nodeIp + ":" + client.config.NodePort, nil
	default:
		return "tcp://" + nodeIp + ":" + client.config.NodePort, nil
	}
}
