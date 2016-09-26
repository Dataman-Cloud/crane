package dockerclient

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strconv"

	"github.com/Dataman-Cloud/crane/src/model"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/swarm"
	"golang.org/x/net/context"
)

const (
	defaultListenAddr = "0.0.0.0:2377"

	flagUpdateRole         = "role"
	flagUpdateAvailability = "availability"
	flagLabelAdd           = "label-add"
	flagLabelRemove        = "label-rm"
	flagLabelUpdate        = "label-update"
	flagEndpointUpdate     = "endpoint-update"
)

// NodeList returns the list of nodes.
func (client *CraneDockerClient) ListNode(opts types.NodeListOptions) ([]swarm.Node, error) {
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

// create a new node
func (client *CraneDockerClient) CreateNode(joiningNode model.JoiningNode) error {
	if joiningNode.Role != swarm.NodeRoleWorker && joiningNode.Role != swarm.NodeRoleManager {
		errMsg := fmt.Sprintf("node role only support %s/%s but got %s",
			swarm.NodeRoleWorker, swarm.NodeRoleManager, joiningNode.Role)
		return cranerror.NewError(CodeErrorNodeRole, errMsg)
	}

	nodeUrl, err := parseEndpoint(joiningNode.Endpoint)
	if err != nil {
		return &cranerror.CraneError{
			Code: CodeGetNodeEndpointError,
			Err:  &cranerror.NodeConnError{ID: "", Err: fmt.Errorf("get node endpoint failed: %s", err.Error())},
		}
	}

	swarmInfo, err := client.InspectSwarm()
	if err != nil {
		return err
	}
	managerInfo, err := client.ManagerInfo()
	if err != nil {
		return err
	}
	advertiseAddr, err := getAdvertiseAddrByEndpoint(joiningNode.Endpoint)
	if err != nil {
		return &cranerror.CraneError{
			Code: CodeGetNodeAdvertiseAddrError,
			Err:  &cranerror.NodeConnError{ID: "", Err: fmt.Errorf("get node advertiseAddr failed: %s", err.Error())},
		}
	}
	var joinToken string
	if joiningNode.Role == swarm.NodeRoleWorker {
		joinToken = swarmInfo.JoinTokens.Worker
	} else if joiningNode.Role == swarm.NodeRoleManager {
		joinToken = swarmInfo.JoinTokens.Manager
	}
	req := swarm.JoinRequest{
		JoinToken:     joinToken,
		AdvertiseAddr: advertiseAddr,
		ListenAddr:    defaultListenAddr,
		RemoteAddrs:   []string{managerInfo.Swarm.NodeAddr},
	}
	opts := docker.JoinSwarmOptions{
		JoinRequest: req,
	}
	// swarm join
	nodeClient, err := client.createNodeClient(nodeUrl)
	if err != nil {
		return err
	}
	if err := nodeClient.JoinSwarm(opts); err != nil {
		return &cranerror.CraneError{
			Code: CodeJoinNodeError,
			Err:  fmt.Errorf("node %s (%s) failed to join cluster: %s", joiningNode.Endpoint, nodeUrl, err.Error()),
		}
	}

	// Store the endpoint in the node label
	nodeId, err := client.getNodeIdByUrl(nodeUrl)
	if err != nil {
		return &cranerror.CraneError{
			Code: CodeVerifyNodeEnpointFailed,
			Err:  &cranerror.NodeConnError{ID: nodeId, Endpoint: joiningNode.Endpoint, Err: fmt.Errorf("verify endpoint failed: %s", err.Error())},
		}
	}
	node, err := client.InspectNode(nodeId)
	if err != nil {
		return err
	}
	if node.Spec.Annotations.Labels == nil {
		node.Spec.Annotations.Labels = make(map[string]string)
	}
	node.Spec.Annotations.Labels[labelNodeEndpoint] = joiningNode.Endpoint
	query := url.Values{}
	query.Set("version", strconv.FormatUint(node.Version.Index, 10))
	_, err = client.sharedHttpClient.POST(nil, client.swarmManagerHttpEndpoint+"/"+path.Join("nodes", nodeId, "update"), query, node.Spec, nil)
	if err != nil {
		return err
	}

	return nil
}

// Inspect node returns the single node.
func (client *CraneDockerClient) InspectNode(nodeId string) (swarm.Node, error) {
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
func (client *CraneDockerClient) RemoveNode(nodeId string) error {
	_, err := client.sharedHttpClient.DELETE(nil, client.swarmManagerHttpEndpoint+"/"+path.Join("nodes", nodeId), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// Update a single node
func (client *CraneDockerClient) UpdateNode(nodeId string, opts model.UpdateOptions) error {
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
	case flagEndpointUpdate:
		if err := client.nodeUpdateEndpoint(nodeId, spec, opts.Options); err != nil {
			return err
		}
	default:
		return cranerror.NewError(CodeErrorUpdateNodeMethod, fmt.Sprintf("Invalid update node method %s", opts.Method))
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
		err = cranerror.NewError(CodeErrorNodeRole, errMsg)
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
		err = cranerror.NewError(CodeErrorNodeAvailability, errMsg)
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

func (client *CraneDockerClient) nodeUpdateEndpoint(nodeId string, spec *swarm.NodeSpec, rawMessage []byte) error {
	var endpoint string
	if err := json.Unmarshal(rawMessage, &endpoint); err != nil {
		return err
	}

	nodeUrl, err := parseEndpoint(endpoint)
	if err != nil {
		return &cranerror.CraneError{
			Code: CodeGetNodeEndpointError,
			Err:  &cranerror.NodeConnError{ID: nodeId, Err: fmt.Errorf("update endpoint failed: %s", err.Error())},
		}
	}

	if err := client.VerifyNodeEndpoint(nodeId, nodeUrl); err != nil {
		return err
	}

	if spec.Annotations.Labels == nil {
		spec.Annotations.Labels = make(map[string]string)
	}

	spec.Annotations.Labels[labelNodeEndpoint] = endpoint
	return nil
}

// docker info
func (client *CraneDockerClient) Info(ctx context.Context) (*docker.DockerInfo, error) {
	swarmNode, err := client.SwarmNode(ctx)
	if err != nil {
		return nil, err
	}
	return swarmNode.Info()
}

func (client *CraneDockerClient) GetDaemonUrlById(nodeId string) (*url.URL, error) {
	node, err := client.InspectNode(nodeId)
	var nodeConnErr *cranerror.NodeConnError
	if err != nil {
		nodeConnErr = &cranerror.NodeConnError{ID: nodeId, Endpoint: "", Err: err}
		return nil, &cranerror.CraneError{Code: CodeGetNodeEndpointError, Err: nodeConnErr}
	}

	endpoint, ok := node.Spec.Annotations.Labels[labelNodeEndpoint]
	if !ok {
		nodeConnErr = &cranerror.NodeConnError{ID: nodeId, Endpoint: endpoint, Err: err}
		return nil, &cranerror.CraneError{Code: CodeGetNodeEndpointError, Err: nodeConnErr}
	}

	return parseEndpoint(endpoint)
}

func (client *CraneDockerClient) getNodeIdByUrl(nodeUrl *url.URL) (string, error) {
	//TODO hardcode may have better way
	if nodeUrl.Scheme == "tcp" {
		nodeUrl.Scheme = "http"
	}

	endpoint := nodeUrl.String()
	content, err := client.sharedHttpClient.GET(nil, endpoint+"/info", url.Values{}, nil)
	if err != nil {
		return "", err
	}

	var nodeInfo types.Info
	if err := json.Unmarshal(content, &nodeInfo); err != nil {
		return "", err
	}

	return nodeInfo.Swarm.NodeID, nil
}
