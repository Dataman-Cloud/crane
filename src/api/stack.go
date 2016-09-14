package api

import (
	"encoding/json"
	"strconv"

	"github.com/Dataman-Cloud/crane/src/dockerclient/model"
	"github.com/Dataman-Cloud/crane/src/plugins/auth"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/httpresponse"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/gin-gonic/gin"
)

const (
	//Stack error code
	CodeCreateStackParamError = "400-11501"
	CodeInvalidStackName      = "503-11502"
	CodeStackNotFound         = "404-11503"

	CodeInvalidGroupId = "400-12001"
)

func (api *Api) UpdateStack(ctx *gin.Context) {}

func (api *Api) CreateStack(ctx *gin.Context) {
	var stackBundle model.Bundle

	if err := ctx.BindJSON(&stackBundle); err != nil {
		switch jsonErr := err.(type) {
		case *json.SyntaxError:
			log.Errorf("Stack JSON syntax error at byte %v: %s", jsonErr.Offset, jsonErr.Error())
		case *json.UnmarshalTypeError:
			log.Errorf("Unexpected type at by type %v. Expected %s but received %s.",
				jsonErr.Offset, jsonErr.Type, jsonErr.Value)
		}

		craneError := cranerror.NewError(CodeCreateStackParamError, err.Error())
		httpresponse.Error(ctx, craneError)
		return
	}

	if api.Config.FeatureEnabled("account") {
		groupId := ctx.DefaultQuery("group_id", "-1")
		groupId = "1"
		gId, err := strconv.ParseUint(groupId, 10, 64)
		if err != nil || gId < 0 {
			log.Error("CreateStack invalid group_id")
			craneError := cranerror.NewError(CodeInvalidGroupId, "invalid group id")
			httpresponse.Error(ctx, craneError)
			return
		}

		perms := auth.PermissionGrantLabelsPairFromGroupIdAndPerm(gId, auth.PermAdmin.Display)
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
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, "success")
	return
}

func (api *Api) ListStack(ctx *gin.Context) {
	stacks, err := api.GetDockerClient().ListStack()
	if err != nil {
		log.Error("Stack deploy got error: ", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, stacks)
	return
}

func (api *Api) InspectStack(ctx *gin.Context) {
	namespace := ctx.Param("namespace")

	bundle, err := api.GetDockerClient().InspectStack(namespace)
	if err != nil {
		log.Error("InspectStack got error: ", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, bundle)
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
		log.Error("ListStackService got error: ", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, servicesStatus)
	return
}

func (api *Api) RemoveStack(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	if err := api.GetDockerClient().RemoveStack(namespace); err != nil {
		log.Error("Remove stack got error: ", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, "success")
	return
}
