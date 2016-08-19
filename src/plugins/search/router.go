package search

import (
	"time"

	docker "github.com/Dataman-Cloud/go-dockerclient"
	"github.com/Dataman-Cloud/rolex/src/dockerclient"
	"github.com/Dataman-Cloud/rolex/src/util/config"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/gin-gonic/gin"
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

type SearchApi struct {
	Config            *config.Config
	RolexDockerClient *dockerclient.RolexDockerClient
}

var searchClient *SearchClient

func (searchApi *SearchApi) RegisterApiForSearch(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	searchV1 := router.Group("/search/v1", middlewares...)
	{
		searchV1.GET("/luckysearch", searchApi.Search)
	}
}

func (searchApi *SearchApi) IndexData() {
	searchClient = &SearchClient{}
	defer func() {
		if err := recover(); err != nil {
			searchApi.IndexData()
		}
	}()
	for {
		searchApi.loadData()
		time.Sleep(time.Minute * time.Duration(searchApi.Config.SearchLoadDataInterval))
	}
}

func (searchApi *SearchApi) loadData() {
	searchClient.Index = []string{}
	searchClient.Store = map[string]Document{}

	if nodes, err := searchApi.RolexDockerClient.ListNode(types.NodeListOptions{}); err == nil {
		for _, node := range nodes {
			searchClient.StoreData(node.ID, Document{
				ID:   node.ID,
				Name: node.Description.Hostname,
				Type: DOCUMENT_NODE,
				Param: map[string]string{
					"NodeId": node.ID,
				},
			})

			searchClient.StoreData(node.Description.Hostname, Document{
				ID:   node.ID,
				Name: node.Description.Hostname,
				Type: DOCUMENT_NODE,
				Param: map[string]string{
					"NodeId": node.ID,
				},
			})

			backContext := context.WithValue(context.Background(), "node_id", node.ID)
			if networks, err := searchApi.
				RolexDockerClient.
				ListNodeNetworks(backContext, docker.NetworkFilterOpts{}); err == nil {
				for _, network := range networks {
					searchClient.StoreData(network.ID, Document{
						Name: network.Name,
						ID:   network.ID,
						Type: DOCUMENT_NETWORK,
						Param: map[string]string{
							"NodeId":    node.ID,
							"NetworkID": network.ID,
						},
					})

					searchClient.StoreData(network.Name, Document{
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
				log.Errorf("get network error: %v", err)
			}

			if volumes, err := searchApi.
				RolexDockerClient.
				ListVolumes(backContext, docker.ListVolumesOptions{}); err == nil {
				for _, volume := range volumes {
					searchClient.StoreData(volume.Name, Document{
						Name: volume.Name,
						Type: DOCUMENT_VOLUME,
						Param: map[string]string{
							"NodeId": node.ID,
						},
					})
				}
			} else {
				log.Errorf("get volume error: %v", err)
			}
		}
	} else {
		log.Errorf("get node list error: %v", err)
	}

	if stacks, err := searchApi.RolexDockerClient.ListStack(); err == nil {
		for _, stack := range stacks {
			//groupId, _ := searchApi.RolexDockerClient.GetStackGroup(stack.Namespace)
			groupId := uint64(1)

			searchClient.StoreData(stack.Namespace, Document{
				ID:      stack.Namespace,
				Type:    DOCUMENT_STACK,
				GroupId: groupId,
				Param: map[string]string{
					"NameSpace": stack.Namespace,
				},
			})

			if services, err := searchApi.
				RolexDockerClient.
				ListStackService(stack.Namespace, types.ServiceListOptions{}); err == nil {
				for _, service := range services {
					searchClient.StoreData(service.ID,
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

					searchClient.StoreData(service.Name,
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

					if tasks, err := searchApi.
						RolexDockerClient.
						ListTasks(types.TaskListOptions{}); err == nil {
						for _, task := range tasks {
							searchClient.StoreData(task.ID,
								Document{
									ID:      task.ID,
									Type:    DOCUMENT_TASK,
									GroupId: groupId,
									Param: map[string]string{
										"NodeId":      task.NodeID,
										"ContainerId": task.Status.ContainerStatus.ContainerID,
									},
								})
						}
					} else {
						log.Errorf("get task list error: %v", err)
					}
				}
			} else {
				log.Errorf("get service error: %v", err)
			}

		}
	} else {
		log.Errorf("get stack list error: %v", err)
	}
}
