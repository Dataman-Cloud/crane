package search

import (
	"time"

	"github.com/Dataman-Cloud/rolex/src/dockerclient"
	"github.com/Dataman-Cloud/rolex/src/util/config"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	goclient "github.com/fsouza/go-dockerclient"
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
	Index             []string
	Store             map[string]Document
}

type Document struct {
	ID    string
	Name  string
	Type  string
	Param map[string]string
}

func (searchApi *SearchApi) RegisterApiForSearch(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	searchV1 := router.Group("/search/v1", middlewares...)
	{
		searchV1.GET("/luckysearch", searchApi.Search)
	}
}

func (searchApi *SearchApi) IndexData() {
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
	searchApi.Index = []string{}
	searchApi.Store = map[string]Document{}

	if nodes, err := searchApi.RolexDockerClient.ListNode(types.NodeListOptions{}); err == nil {
		for _, node := range nodes {
			searchApi.Index = append(searchApi.Index, node.ID)
			searchApi.Store[node.ID] = Document{
				ID:   node.ID,
				Name: node.Description.Hostname,
				Type: DOCUMENT_NODE,
				Param: map[string]string{
					"NodeId": node.ID,
				},
			}

			searchApi.Index = append(searchApi.Index, node.Description.Hostname)
			searchApi.Store[node.ID] = Document{
				ID:   node.ID,
				Name: node.Description.Hostname,
				Type: DOCUMENT_NODE,
				Param: map[string]string{
					"NodeId": node.ID,
				},
			}

			backContext := context.WithValue(context.Background(), "node_id", node.ID)
			if networks, err := searchApi.
				RolexDockerClient.
				ListNodeNetworks(backContext, goclient.NetworkFilterOpts{}); err == nil {
				for _, network := range networks {
					searchApi.Index = append(searchApi.Index, network.ID)
					searchApi.Store[network.ID] = Document{
						Name: network.Name,
						ID:   network.ID,
						Type: DOCUMENT_NETWORK,
						Param: map[string]string{
							"NodeId":    node.ID,
							"NetworkID": network.ID,
						},
					}

					searchApi.Index = append(searchApi.Index, network.Name)
					searchApi.Store[network.Name] = Document{
						Name: network.Name,
						ID:   network.ID,
						Type: DOCUMENT_NETWORK,
						Param: map[string]string{
							"NodeId":    node.ID,
							"NetworkID": network.ID,
						},
					}
				}
			} else {
				log.Errorf("get network error: %v", err)
			}

			if volumes, err := searchApi.
				RolexDockerClient.
				ListVolumes(backContext, goclient.ListVolumesOptions{}); err == nil {
				for _, volume := range volumes {
					searchApi.Index = append(searchApi.Index, volume.Name)
					searchApi.Store[volume.Name] = Document{
						Name: volume.Name,
						Type: DOCUMENT_VOLUME,
						Param: map[string]string{
							"NodeId": node.ID,
						},
					}
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
			searchApi.Index = append(searchApi.Index, stack.Namespace)
			searchApi.Store[stack.Namespace] = Document{
				ID:   stack.Namespace,
				Type: DOCUMENT_STACK,
				Param: map[string]string{
					"NameSpace": stack.Namespace,
				},
			}

			if services, err := searchApi.
				RolexDockerClient.
				ListStackService(stack.Namespace, types.ServiceListOptions{}); err == nil {
				for _, service := range services {
					searchApi.Index =
						append(searchApi.Index, service.ID)
					searchApi.Store[service.ID] =
						Document{
							ID:   service.ID,
							Name: stack.Namespace,
							Type: DOCUMENT_SERVICE,
							Param: map[string]string{
								"NameSpace": stack.Namespace,
								"ServiceId": service.ID,
							},
						}

					searchApi.Index =
						append(searchApi.Index, service.Name)
					searchApi.Store[service.Name] =
						Document{
							ID:   service.ID,
							Name: stack.Namespace,
							Type: DOCUMENT_SERVICE,
							Param: map[string]string{
								"NameSpace": stack.Namespace,
								"ServiceId": service.ID,
							},
						}

					if tasks, err := searchApi.
						RolexDockerClient.
						ListTasks(types.TaskListOptions{}); err == nil {
						for _, task := range tasks {
							searchApi.Index =
								append(searchApi.Index, task.ID)
							searchApi.Store[task.ID] =
								Document{
									ID:   task.ID,
									Type: DOCUMENT_TASK,
									Param: map[string]string{
										"NodeId":      task.NodeID,
										"ContainerId": task.Status.ContainerStatus.ContainerID,
									},
								}
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
