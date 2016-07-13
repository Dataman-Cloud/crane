package api

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"github.com/Dataman-Cloud/rolex/dockerclient"
)

func (api *Api) UpdateStack(ctx *gin.Context) {}
func (api *Api) RemoveStack(ctx *gin.Context) {}

func (api *Api) CreateStack(ctx *gin.Context) {
	stackBundle := dockerclient.Bundle{}

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

	if err := api.GetDockerClient().DeployStack(&stackBundle); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.Error("Stack deploy got error: ", err)
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

	servicesStatus, err := api.GetDockerClient().ListStackService(namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.Error("ListStackService got error: ", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": servicesStatus})
	return

}
