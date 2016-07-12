package api

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/client/bundlefile"
	"github.com/gin-gonic/gin"
)

func (api *Api) InspectStack(ctx *gin.Context) {}
func (api *Api) ListStack(ctx *gin.Context)    {}
func (api *Api) UpdateStack(ctx *gin.Context)  {}
func (api *Api) RemoveStack(ctx *gin.Context)  {}

func (api *Api) StackCreate(ctx *gin.Context) {
	stackBundlefile := bundlefile.Bundlefile{}

	if err := ctx.BindJSON(&stackBundlefile); err != nil {
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

	namespace := ctx.Param("name")
	if err := api.GetDockerClient().StackDeploy(&stackBundlefile, namespace); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.Error("Stack deploy got error: ", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "success"})
}
