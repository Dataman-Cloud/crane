package plugins

import (
	"github.com/Dataman-Cloud/go-component/search"
	"github.com/Dataman-Cloud/rolex/src/dockerclient"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
)

const (
	DOCUMENT_NODE    = "node"
	DOCUMENT_STACK   = "stack"
	DOCUMENT_SERVICE = "service"
	DOCUMENT_TASK    = "task"
	DOCUMENT_NETWORK = "network"
	DOCUMENT_VOLUME  = "volume"
)

type RolexIndexer struct {
	search.Indexer

	RolexDockerClient *dockerclient.RolexDockerClient
}

func NewRolexIndex(RolexDockerClient *dockerclient.RolexDockerClient) *RolexIndexer {
	return &RolexIndexer{RolexDockerClient: RolexDockerClient}
}

func (indexer *RolexIndexer) Index(store *search.DocumentStorage) {
	if nodes, err := indexer.RolexDockerClient.ListNode(types.NodeListOptions{}); err == nil {
		for _, node := range nodes {
			store.Set(node.ID+node.Description.Hostname, search.Document{
				ID:   node.ID,
				Name: node.Description.Hostname,
				Type: DOCUMENT_NODE,
				Param: map[string]string{
					"NodeId": node.ID,
				},
			})

			backContext := context.WithValue(context.Background(), "node_id", node.ID)
			if networks, err := indexer.
				RolexDockerClient.
				ListNodeNetworks(backContext, docker.NetworkFilterOpts{}); err == nil {
				for _, network := range networks {
					store.Set(network.ID+network.Name+node.ID, search.Document{
						Name: network.Name,
						ID:   network.ID,
						Type: DOCUMENT_NETWORK,
						Param: map[string]string{
							"NodeId":    node.ID,
							"NetworkID": network.ID,
						},
					})

				}
			} else {
				log.Warnf("get network error: %v", err)
			}

			if volumes, err := indexer.
				RolexDockerClient.
				ListVolumes(backContext, docker.ListVolumesOptions{}); err == nil {
				for _, volume := range volumes {
					store.Set(volume.Name, search.Document{
						Name: volume.Name,
						Type: DOCUMENT_VOLUME,
						Param: map[string]string{
							"NodeId": node.ID,
						},
					})
				}
			} else {
				log.Warnf("get volume error: %v", err)
			}
		}
	} else {
		log.Warnf("get node list error: %v", err)
	}

	if stacks, err := indexer.RolexDockerClient.ListStack(); err == nil {
		for _, stack := range stacks {
			//groupId, _ := indexer.RolexDockerClient.GetStackGroup(stack.Namespace)
			groupId := uint64(1)

			store.Set(stack.Namespace, search.Document{
				ID:      stack.Namespace,
				Type:    DOCUMENT_STACK,
				GroupId: groupId,
				Param: map[string]string{
					"NameSpace": stack.Namespace,
				},
			})

			if services, err := indexer.
				RolexDockerClient.
				ListStackService(stack.Namespace, types.ServiceListOptions{}); err == nil {
				for _, service := range services {
					store.Set(service.ID+stack.Namespace,
						search.Document{
							ID:      service.ID,
							Name:    stack.Namespace,
							Type:    DOCUMENT_SERVICE,
							GroupId: groupId,
							Param: map[string]string{
								"NameSpace": stack.Namespace,
								"ServiceId": service.ID,
							},
						})

					if tasks, err := indexer.
						RolexDockerClient.
						ListTasks(types.TaskListOptions{}); err == nil {
						for _, task := range tasks {
							store.Set(task.ID,
								search.Document{
									ID:      task.ID,
									Type:    DOCUMENT_TASK,
									GroupId: groupId,
									Param: map[string]string{
										"NodeId":      task.NodeID,
										"ContainerId": task.Status.ContainerStatus.ContainerID,
									},
								})
							store.Set(task.Status.ContainerStatus.ContainerID,
								search.Document{
									ID:      task.Status.ContainerStatus.ContainerID,
									Type:    DOCUMENT_TASK,
									GroupId: groupId,
									Param: map[string]string{
										"NodeId":      task.NodeID,
										"ContainerId": task.Status.ContainerStatus.ContainerID,
									},
								})
						}
					} else {
						log.Warnf("get task list error: %v", err)
					}
				}
			} else {
				log.Warnf("get service error: %v", err)
			}

		}
	} else {
		log.Warnf("get stack list error: %v", err)
	}
}
