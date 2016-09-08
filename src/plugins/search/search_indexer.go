package search

import (
	"github.com/Dataman-Cloud/crane/src/dockerclient"

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

type CraneIndexer struct {
	Indexer

	CraneDockerClient *dockerclient.CraneDockerClient
}

func NewCraneIndex(CraneDockerClient *dockerclient.CraneDockerClient) *CraneIndexer {
	return &CraneIndexer{CraneDockerClient: CraneDockerClient}
}

func (indexer *CraneIndexer) Index(store *DocumentStorage) {
	if nodes, err := indexer.CraneDockerClient.ListNode(types.NodeListOptions{}); err == nil {
		for _, node := range nodes {
			store.Set(node.ID+node.Description.Hostname, Document{
				ID:   node.ID,
				Name: node.Description.Hostname,
				Type: DOCUMENT_NODE,
				Param: map[string]string{
					"NodeId": node.ID,
				},
			})

			backContext := context.WithValue(context.Background(), "node_id", node.ID)
			if networks, err := indexer.
				CraneDockerClient.
				ListNodeNetworks(backContext, docker.NetworkFilterOpts{}); err == nil {
				for _, network := range networks {
					store.Set(network.ID+network.Name+node.ID, Document{
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
				CraneDockerClient.
				ListVolumes(backContext, docker.ListVolumesOptions{}); err == nil {
				for _, volume := range volumes {
					store.Set(volume.Name, Document{
						Name: volume.Name,
						Type: DOCUMENT_VOLUME,
						Param: map[string]string{
							"NodeId":     node.ID,
							"VolumeName": volume.Name,
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

	if stacks, err := indexer.CraneDockerClient.ListStack(); err == nil {
		for _, stack := range stacks {
			//groupId, _ := indexer.CraneDockerClient.GetStackGroup(stack.Namespace)
			groupId := uint64(1)

			store.Set(stack.Namespace, Document{
				ID:      stack.Namespace,
				Type:    DOCUMENT_STACK,
				GroupId: groupId,
				Param: map[string]string{
					"NameSpace": stack.Namespace,
				},
			})

			if services, err := indexer.
				CraneDockerClient.
				ListStackService(stack.Namespace, types.ServiceListOptions{}); err == nil {
				for _, service := range services {
					store.Set(service.ID+stack.Namespace,
						Document{
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
						CraneDockerClient.
						ListTasks(types.TaskListOptions{}); err == nil {
						for _, task := range tasks {
							store.Set(task.ID,
								Document{
									ID:      task.ID,
									Type:    DOCUMENT_TASK,
									GroupId: groupId,
									Param: map[string]string{
										"NodeId":      task.NodeID,
										"ContainerId": task.Status.ContainerStatus.ContainerID,
									},
								})
							store.Set(task.Status.ContainerStatus.ContainerID,
								Document{
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
