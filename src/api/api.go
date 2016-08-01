package api

import (
	"net/http"

	"github.com/Dataman-Cloud/rolex/src/dockerclient"
	"github.com/Dataman-Cloud/rolex/src/util/config"
	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"

	"github.com/gin-gonic/gin"
)

type Api struct {
	Client *dockerclient.RolexDockerClient
	Config *config.Config
}

func (api *Api) GetDockerClient() *dockerclient.RolexDockerClient {
	return api.Client
}

func (api *Api) GetConfig() *config.Config {
	return api.Config
}

// RHttprespnse retrun none error code 200
func (api *Api) HttpResponse(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"code": rolexerror.CodeOk, "data": data})
		return
	}

	api.HttpErrorResponse(ctx, err)
	return
}

// RHttprespnse retrun none error code 201
func (api *Api) HttpCreateResponse(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		ctx.JSON(http.StatusCreated, gin.H{"code": rolexerror.CodeOk, "data": data})
		return
	}

	api.HttpErrorResponse(ctx, err)
	return
}

// RHttprespnse retrun none error code 204
func (api *Api) HttpDeleteResponse(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		ctx.JSON(http.StatusNoContent, gin.H{"code": rolexerror.CodeOk, "data": data})
		return
	}

	api.HttpErrorResponse(ctx, err)
	return
}

// RHttprespnse retrun none error code 202
func (api *Api) HttpUpdateResponse(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		ctx.JSON(http.StatusAccepted, gin.H{"code": rolexerror.CodeOk, "data": data})
		return
	}

	api.HttpErrorResponse(ctx, err)
	return
}

func (api *Api) HttpErrorResponse(ctx *gin.Context, err error) {
	rerror, ok := err.(rolexerror.RolexError)
	if !ok {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"code": rolexerror.CodeUndefined, "data": err})
		return
	}

	var httpCode int
	switch rerror.Code {
	case rolexerror.CodeNetworkPredefined:
		httpCode = http.StatusForbidden
	default:
		httpCode = http.StatusServiceUnavailable
	}

	ctx.JSON(httpCode, gin.H{"code": rerror.Code, "data": err})
	return
}
