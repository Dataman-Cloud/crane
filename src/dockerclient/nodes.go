package dockerclient

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/Dataman-Cloud/go-component/utils/dmerror"
	"github.com/Dataman-Cloud/rolex/src/model"
	"github.com/Dataman-Cloud/rolex/src/util/config"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"golang.org/x/net/context"
)

const (
	flagUpdateRole         = "role"
	flagUpdateAvailability = "availability"
	flagLabelAdd           = "label-add"
	flagLabelRemove        = "label-rm"
	flagLabelUpdate        = "label-update"
)

const (
	CodeErrorUpdateNodeMethod = "503-11302"
	CodeErrorNodeRole         = "503-11303"
	CodeErrorNodeAvailability = "503-11304"
	CodeGetNodeInfoError      = "503-11305"
)

const (
	labelNodeEndpoint = "dm.swarm.node.endpoint"
)

// NodeList returns the list of nodes.
func (client *RolexDockerClient) ListNode(opts types.NodeListOptions) ([]swarm.Node, error) {
	var nodes []swarm.Node

	content, err := client.sharedHttpClient.GET(nil, client.swarmManagerHttpEndpoint+"/nodes", nil, nil)
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

	content, err := client.sharedHttpClient.GET(nil, client.swarmManagerHttpEndpoint+"/"+path.Join("nodes", nodeId), nil, nil)
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
	_, err := client.sharedHttpClient.DELETE(nil, client.swarmManagerHttpEndpoint+"/"+path.Join("nodes", nodeId), nil, nil)
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
	case flagLabelUpdate:
		if err := nodeUpdateLabels(spec, opts.Options); err != nil {
			return err
		}
	default:
		return dmerror.NewError(CodeErrorUpdateNodeMethod, fmt.Sprintf("Invalid update node method %s", opts.Method))
	}

	query := url.Values{}
	query.Set("version", strconv.FormatUint(node.Version.Index, 10))
	_, err = client.sharedHttpClient.POST(nil, client.swarmManagerHttpEndpoint+"/"+path.Join("nodes", nodeId, "update"), query, node.Spec, nil)
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
		err = dmerror.NewError(CodeErrorNodeRole, errMsg)
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
		err = dmerror.NewError(CodeErrorNodeAvailability, errMsg)
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

func nodeUpdateLabels(spec *swarm.NodeSpec, rawMessage []byte) error {
	var labelsUpdate map[string]string
	if err := json.Unmarshal(rawMessage, &labelsUpdate); err != nil {
		return err
	}

	spec.Annotations.Labels = labelsUpdate

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
func (client *RolexDockerClient) Info(ctx context.Context) (*docker.DockerInfo, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.Info()
}

func (client *RolexDockerClient) NodeDaemonUrl(nodeId string) (*url.URL, error) {
	node, err := client.InspectNode(nodeId)
	var nodeConnErr *rolexerror.NodeConnError
	if err != nil {
		nodeConnErr = &rolexerror.NodeConnError{ID: nodeId, Endpoint: "", Err: err}
		return nil, &dmerror.DmError{Code: CodeGetNodeEndpointError, Err: nodeConnErr}
	}

	endpoint, ok := node.Spec.Annotations.Labels[labelNodeEndpoint]
	if !ok {
		nodeConnErr = &rolexerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: err}
		return nil, &dmerror.DmError{Code: CodeGetNodeEndpointError, Err: nodeConnErr}
	}

	return parseEndpoint(endpoint)
}

func parseEndpoint(endpoint string) (*url.URL, error) {
	conf := config.GetConfig()
	if !strings.Contains(endpoint, "://") {
		endpoint = conf.DockerEntryScheme + "://" + endpoint
	}

	u, err := url.Parse(endpoint)

	if err != nil {
		return nil, err
	}

	if !strings.Contains(u.Host, ":") {
		u.Host = u.Host + ":" + conf.DockerEntryPort
	}

	return u, nil
}

func (client *RolexDockerClient) UpdateNodeEndpoint(nodeId, endpoint string) error {
	nodeUrl, err := parseEndpoint(endpoint)
	if err != nil {
		return &dmerror.DmError{
			Code: CodeGetNodeEndpointError,
			Err:  &rolexerror.NodeConnError{ID: nodeId, Err: fmt.Errorf("update endpoint failed: %s", err.Error())},
		}
	}

	if err := client.VerifyNodeEndpoint(nodeId, nodeUrl); err != nil {
		return err
	}

	node, err := client.InspectNode(nodeId)
	if err != nil {
		return err
	}

	spec := &node.Spec

	if spec.Annotations.Labels == nil {
		spec.Annotations.Labels = make(map[string]string)
	}

	spec.Annotations.Labels[labelNodeEndpoint] = endpoint
	query := url.Values{}
	query.Set("version", strconv.FormatUint(node.Version.Index, 10))
	_, err = client.sharedHttpClient.POST(nil, client.swarmManagerHttpEndpoint+"/"+path.Join("nodes", nodeId, "update"), query, node.Spec, nil)
	if err != nil {
		return err
	}

	return nil

}
