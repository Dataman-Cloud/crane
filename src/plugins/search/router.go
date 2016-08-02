package search

import (
	"fmt"
	"time"

	"github.com/Dataman-Cloud/rolex/src/dockerclient"
	"github.com/Dataman-Cloud/rolex/src/util/config"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/gin-gonic/gin"
)

const (
	NODE_INFO    = "/node/detail/%s/config"
	NETWORK_INFO = "/network/%s/%s/config"
	STACK_INFO   = "/stack/detail/%s/service"
	SERVICE_INFO = "/stack/serviceDetail/%s/%s/config"
	TASK_INFO    = "/node/containerDetail/%s/%s/config"
)

const (
	METHOD_GET = "get"
)

const (
	DOCUMENT_NODE    = "node"
	DOCUMENT_STACK   = "stack"
	DOCUMENT_SERVICE = "service"
	DOCUMENT_TASK    = "task"
)

type SearchApi struct {
	Config            *config.Config
	RolexDockerClient *dockerclient.RolexDockerClient
	Index             []string
	Store             map[string]Document
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
