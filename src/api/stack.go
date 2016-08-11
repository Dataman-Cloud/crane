package api

import (
	"encoding/json"
	"strconv"

	"github.com/Dataman-Cloud/rolex/src/dockerclient"
	"github.com/Dataman-Cloud/rolex/src/dockerclient/model"
	"github.com/Dataman-Cloud/rolex/src/plugins/auth"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"
	"github.com/Dataman-Cloud/rolex/src/util/rolexgin"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/gin-gonic/gin"
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

		rerror := rolexerror.NewRolexError(rolexerror.CodeCreateStackParamError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rerror)
		return
	}

	if api.Config.FeatureEnabled("account") {
		groupId := ctx.DefaultQuery("group_id", "-1")
		gId, err := strconv.ParseUint(groupId, 10, 64)
		if err != nil || gId < 0 {
			log.Error("CreateStack invalid group_id")
			rerror := rolexerror.NewRolexError(rolexerror.CodeInvalidGroupId, "invalid group id")
			rolexgin.HttpErrorResponse(ctx, rerror)
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
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, "success")
	return
}

func (api *Api) ListStack(ctx *gin.Context) {
	stacks, err := api.GetDockerClient().ListStack()

	if groups, ok := ctx.Get("groups"); ok {
		var astacks []dockerclient.Stack
		for _, stack := range stacks {
			groupId, err := api.GetDockerClient().GetStackGroup(stack.Namespace)
			if err != nil {
				break
			}
			for _, group := range groups.([]auth.Group) {
				if group.ID == groupId {
					astacks = append(astacks, stack)
					break
				}
			}
		}
		rolexgin.HttpOkResponse(ctx, astacks)
		return
	}

	if err != nil {
		log.Error("Stack deploy got error: ", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, stacks)
	return
}

func (api *Api) InspectStack(ctx *gin.Context) {
	namespace := ctx.Param("namespace")

	bundle, err := api.GetDockerClient().InspectStack(namespace)
	if err != nil {
		log.Error("InspectStack got error: ", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, bundle)
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
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, servicesStatus)
	return
}

func (api *Api) RemoveStack(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	if err := api.GetDockerClient().RemoveStack(namespace); err != nil {
		log.Error("Remove stack got error: ", err)
		rolexgin.HttpErrorResponse(ctx, err)
		return
	}

	rolexgin.HttpOkResponse(ctx, "success")
	return
}
