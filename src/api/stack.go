package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Dataman-Cloud/rolex/src/dockerclient/model"
	"github.com/Dataman-Cloud/rolex/src/plugins/auth"
	"github.com/Dataman-Cloud/rolex/src/util"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/gin-gonic/gin"
)

func (api *Api) UpdateStack(ctx *gin.Context) {}

func (api *Api) CreateStack(ctx *gin.Context) {
	stackBundle := model.Bundle{}

	if err := ctx.BindJSON(&stackBundle); err != nil {
		switch jsonErr := err.(type) {
		case *json.SyntaxError:
			log.Errorf("Stack JSON syntax error at byte %v: %s", jsonErr.Offset, jsonErr.Error())
		case *json.UnmarshalTypeError:
			log.Errorf("Unexpected type at by type %v. Expected %s but received %s.",
				jsonErr.Offset, jsonErr.Type, jsonErr.Value)
		}

		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if api.Config.FeatureEnabled("account") {
		if ctx.Query("group_id") == "" {
			log.Error("invalid group_id")
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 0, "data": "invalid group_id"})
			return
		}

		groupId, err := strconv.ParseUint(ctx.Query("group_id"), 10, 64)
		if err != nil {
			log.Error("invalid group_id")
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 0, "data": "invalid group_id"})
			return
		}

		perms := auth.PermissionGrantLabelsPairFromGroupIdAndPerm(groupId, auth.PermAdmin.Display)
		for sk, sv := range stackBundle.Stack.Services {
			if sv.Labels == nil {
				sv.Labels = perms
			} else {
				for pk, pv := range perms {
					sv.Labels[pk] = pv
				}
			}
			stackBundle.Stack.Services[sk] = sv
		}
	}

	if err := api.GetDockerClient().DeployStack(&stackBundle); err != nil {
		log.Error("Stack deploy got error: ", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "success"})
	return
}

func (api *Api) ListStack(ctx *gin.Context) {
	stacks, err := api.GetDockerClient().ListStack()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.Error("Stack deploy got error: ", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": stacks})
	return
}

func (api *Api) InspectStack(ctx *gin.Context) {
	namespace := ctx.Param("namespace")

	bundle, err := api.GetDockerClient().InspectStack(namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.Error("InspectStack got error: ", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": bundle})
	return
}

func (api *Api) ListStackService(ctx *gin.Context) {
	namespace := ctx.Param("namespace")

	opts := types.ServiceListOptions{}
	if labelFilters_, found := ctx.Get("labelFilters"); found {
		labelFilters := labelFilters_.(map[string]string)
		args := filters.NewArgs()
		for k, _ := range labelFilters {
			args.Add("label", k)
		}
		opts.Filter = args
	}

	servicesStatus, err := api.GetDockerClient().ListStackService(namespace, opts)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.Error("ListStackService got error: ", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": servicesStatus})
	return
}

func (api *Api) RemoveStack(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	if err := api.GetDockerClient().RemoveStack(namespace); err != nil {
		log.Error("Remove stack got error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": util.PARAMETER_ERROR, "data": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "removed" + namespace + "success"})
	return
}
