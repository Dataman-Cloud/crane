package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Dataman-Cloud/rolex/dockerclient"
	"github.com/Dataman-Cloud/rolex/util/config"
	"github.com/Dataman-Cloud/rolex/util/rerror"

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

func (api *Api) OK(ctx *gin.Context, httpCode int, data interface{}) {
	ctx.JSON(httpCode, gin.H{"code": 0, "data": data})
}

func (api *Api) ERROR(ctx *gin.Context, err error) {
	var e rerror.RolexError
	var ok bool
	if e, ok = err.(rerror.RolexError); !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": rerror.DEFAULT_ERROR_CODE, "data": err.Error()})
		return
	}

	codes := strings.SplitN(e.ErrorCode, "-", 2)
	var httpStatusCode int
	if len(codes) != 2 {
		httpStatusCode = http.StatusInternalServerError
	} else {
		httpStatusCode, _ = strconv.Atoi(codes[1])
	}
	ctx.JSON(httpStatusCode, gin.H{"code": e.ErrorCode, "data": e.ErrorMessage})
}
