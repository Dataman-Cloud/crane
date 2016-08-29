package dmgin

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Dataman-Cloud/go-component/utils/dmerror"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

const (
	CodeOk        = 0
	CodeUndefined = 10001
)

// RHttprespnse retrun none error code 200
func HttpOkResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": CodeOk, "data": data})
	return
}

// RHttprespnse retrun none error code 201
func HttpCreateResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{"code": CodeOk, "data": data})
	return
}

// RHttprespnse retrun none error code 204
func HttpDeleteResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusNoContent, gin.H{"code": CodeOk, "data": data})
	return
}

// RHttprespnse retrun none error code 202
func HttpUpdateResponse(ctx *gin.Context, err error, data interface{}) {
	ctx.JSON(http.StatusAccepted, gin.H{"code": CodeOk, "data": data})
	return
}

func HttpErrorResponse(ctx *gin.Context, err error) {
	log.Errorf("[%s] %s GOT error: %s", ctx.Request.Method, ctx.Request.URL.Path, err.Error())

	rerror, ok := err.(*dmerror.DmError)
	if !ok {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"code": CodeUndefined, "data": err, "message": err.Error(), "source": "docker"})
		return
	}

	httpCode := http.StatusServiceUnavailable
	errCode := CodeUndefined

	codes := strings.Split(rerror.Code, "-")
	if len(codes) == 2 {
		httpCode, _ = strconv.Atoi(codes[0])
		errCode, _ = strconv.Atoi(codes[1])
	}
	ctx.JSON(httpCode, gin.H{"code": errCode, "data": rerror.Err, "message": rerror.Err.Error(), "source": "rolex"})
	return
}
