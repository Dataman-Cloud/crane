package api

import (
	"time"

	"github.com/Dataman-Cloud/crane/src/api/middlewares"
	"github.com/Dataman-Cloud/crane/src/plugins/apiplugin"
	"github.com/Dataman-Cloud/crane/src/plugins/auth"
	authApi "github.com/Dataman-Cloud/crane/src/plugins/auth/api"
	"github.com/Dataman-Cloud/crane/src/plugins/search"
	"github.com/Dataman-Cloud/crane/src/utils/log"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func (api *Api) ApiRouter() *gin.Engine {
	router := gin.New()
	Authorization := middlewares.Authorization
	AuthorizeServiceAccess := middlewares.AuthorizeServiceAccess

	router.Use(log.Ginrus(logrus.StandardLogger(), time.RFC3339, true), gin.Recovery())
	router.Use(middlewares.OptionHandler())
	router.Use(middlewares.CraneApiContext())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "pass")
	})

	v1 := router.Group("/api/v1", Authorization, middlewares.ListIntercept())
	{
		v1.GET("/nodes", api.ListNodes)
		v1.POST("/nodes", api.CreateNode)
		v1.GET("/nodes/:node_id", api.InspectNode)
		v1.GET("/nodes/:node_id/info", api.Info)
		v1.PATCH("/nodes/:node_id", api.UpdateNode)
		v1.DELETE("/nodes/:node_id", api.RemoveNode)
		// Going to delegate to /nodes/:id
		// v1.GET("/nodes/manager_info", api.ManagerInfo)

		// Containers
		v1.GET("/nodes/:node_id/containers/:container_id/terminal", api.ConnectContainer)
		v1.GET("/nodes/:node_id/containers", api.ListContainers)
		v1.GET("/nodes/:node_id/containers/:container_id", api.InspectContainer)
		v1.GET("/nodes/:node_id/containers/:container_id/diff", api.DiffContainer)
		v1.DELETE("/nodes/:node_id/containers/:container_id", api.DeleteContainer)
		v1.GET("/nodes/:node_id/containers/:container_id/logs", api.LogsContainer)
		v1.GET("/nodes/:node_id/containers/:container_id/stats", api.StatsContainer)

		// Images
		v1.GET("/nodes/:node_id/images", api.ListImages)
		v1.GET("/nodes/:node_id/images/:image_id", api.InspectImage)
		v1.GET("/nodes/:node_id/images/:image_id/history", api.ImageHistory)
		v1.DELETE("/nodes/:node_id/images/:image_id", api.RemoveImage)

		// Volumes
		v1.GET("/nodes/:node_id/volumes", api.ListVolume)
		v1.GET("/nodes/:node_id/volumes/:volume_id", api.InspectVolume)
		v1.POST("/nodes/:node_id/volumes", api.CreateVolume)
		v1.DELETE("/nodes/:node_id/volumes/:volume_id", api.RemoveVolume)

		// Networks
		v1.POST("/nodes/:node_id/networks", api.CreateNodeNetwork)
		v1.GET("/nodes/:node_id/networks", api.ListNodeNetworks)
		v1.GET("/nodes/:node_id/networks/:network_id", api.InspectNodeNetwork)
		v1.PATCH("/nodes/:node_id/networks/:network_id", api.ConnectNodeNetwork)

		v1.POST("/networks", api.CreateNetwork)
		v1.GET("/networks", api.ListNetworks)
		v1.DELETE("/networks/:network_id", api.RemoveNetwork)
		v1.GET("/networks/:network_id", api.InspectNetwork)
		v1.PATCH("/networks/:network_id", api.ConnectNetwork)

		v1.POST("/stacks", api.CreateStack)
		v1.GET("/stacks", api.ListStack)
		v1.GET("/stacks/:namespace", api.InspectStack)
		v1.DELETE("/stacks/:namespace", api.RemoveStack)
		v1.PUT("/stacks/:namespace/services/:service_id", api.UpdateService)
		v1.PATCH("/stacks/:namespace/services/:service_id", api.ScaleService)
		v1.GET("/stacks/:namespace/services/:service_id", AuthorizeServiceAccess(auth.PermReadOnly), api.InspectService)
		v1.GET("/stacks/:namespace/services", AuthorizeServiceAccess(auth.PermReadOnly), api.ListStackService)
		v1.GET("/stacks/:namespace/services/:service_id/logs", api.LogsService)
		v1.GET("/stacks/:namespace/services/:service_id/stats", api.StatsService)
		v1.GET("/stacks/:namespace/services/:service_id/tasks", api.ListTasks)
		v1.GET("/stacks/:namespace/services/:service_id/tasks/:task_id", api.InspectTask)
		v1.GET("/stacks/:namespace/services/:service_id/cd_url", api.ServiceCDAddr)
	}

	if plugin, ok := apiplugin.ApiPlugins[apiplugin.Account]; ok {
		accountApi, ok := plugin.Instance.(*authApi.AccountApi)
		if ok {
			accountApi.CraneDockerClient = api.Client
			Authorization = accountApi.Authorization
			accountApi.ApiRegister(router, middlewares.ListIntercept())
		}
	}

	for _, plugin := range apiplugin.ApiPlugins {
		if plugin.Instance != nil {
			switch plugin.Name {
			case apiplugin.Search:
				searchApi, ok := plugin.Instance.(*search.SearchApi)
				if ok {
					searchApi.Indexer = search.NewCraneIndex(api.Client)
					searchApi.ApiRegister(router, Authorization, middlewares.ListIntercept())
				}
			case apiplugin.Account:
			default:
				plugin.Instance.ApiRegister(router, Authorization, middlewares.ListIntercept())
			}
		}
	}

	router.PUT("/api/v1/stacks/:namespace/services/:service_id/rolling_update", api.UpdateServiceImage) // skip authorization, public access

	misc := router.Group("/misc/v1")
	{
		misc.GET("/help", api.Help(router))
		misc.GET("/config", api.CraneConfig)
		misc.GET("/health", api.HealthCheck)
	}

	return router
}
