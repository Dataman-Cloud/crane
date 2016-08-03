package search

import (
	"fmt"
	"sync"
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
	NODE_INFO    = "/node/detail/%s/config"
	NETWORK_INFO = "/network/%s/%s/config"
	STACK_INFO   = "/stack/detail/%s/service"
	SERVICE_INFO = "/stack/serviceDetail/%s/%s/config"
	TASK_INFO    = "/node/containerDetail/%s/%s/config"
	VOLUME_INFO  = "/node/detail/%s/volume"
)

const (
	METHOD_GET = "get"
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
	LoadLock          sync.Mutex
}

type Document struct {
	ID   string
	Name string
	Url  string
	Type string
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
		time.Sleep(time.Minute * time.Duration(searchApi.Config.LoadDataInterval))
	}
}

func (searchApi *SearchApi) loadData() {
	searchApi.LoadLock.Lock()
	defer searchApi.LoadLock.Unlock()

	searchApi.Index = []string{}
	searchApi.Store = map[string]Document{}

	if nodes, err := searchApi.RolexDockerClient.ListNode(types.NodeListOptions{}); err == nil {
		for _, node := range nodes {
			searchApi.Index = append(searchApi.Index, node.ID)
			searchApi.Store[node.ID] = Document{
				ID:   node.ID,
				Url:  fmt.Sprintf(NODE_INFO, node.ID),
				Type: DOCUMENT_NODE,
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
						Url:  fmt.Sprintf(NETWORK_INFO, node.ID, network.ID),
						Type: DOCUMENT_NETWORK,
					}

					searchApi.Index = append(searchApi.Index, network.Name)
					searchApi.Store[network.Name] = Document{
						Name: network.Name,
						ID:   network.ID,
						Url:  fmt.Sprintf(NETWORK_INFO, node.ID, network.ID),
						Type: DOCUMENT_NETWORK,
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
						Url:  fmt.Sprintf(VOLUME_INFO, node.ID),
						Type: DOCUMENT_VOLUME,
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
				Url:  fmt.Sprintf(STACK_INFO, stack.Namespace),
				Type: DOCUMENT_STACK,
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
							Url:  fmt.Sprintf(SERVICE_INFO, stack.Namespace, service.ID),
							Type: DOCUMENT_SERVICE,
						}

					searchApi.Index =
						append(searchApi.Index, service.Name)
					searchApi.Store[service.Name] =
						Document{
							ID:   service.ID,
							Name: stack.Namespace,
							Url:  fmt.Sprintf(SERVICE_INFO, stack.Namespace, service.ID),
							Type: DOCUMENT_SERVICE,
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
									Url:  fmt.Sprintf(TASK_INFO, task.NodeID, task.Status.ContainerStatus.ContainerID),
									Type: DOCUMENT_TASK,
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
