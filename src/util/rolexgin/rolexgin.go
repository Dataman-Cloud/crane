package rolexgin

import (
	"net/http"

	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"

	"github.com/gin-gonic/gin"
)

// RHttprespnse retrun none error code 200
func HttpOkResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": rolexerror.CodeOk, "data": data})
	return
}

// RHttprespnse retrun none error code 201
func HttpCreateResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{"code": rolexerror.CodeOk, "data": data})
	return
}

// RHttprespnse retrun none error code 204
func HttpDeleteResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusNoContent, gin.H{"code": rolexerror.CodeOk, "data": data})
	return
}

// RHttprespnse retrun none error code 202
func HttpUpdateResponse(ctx *gin.Context, err error, data interface{}) {
	ctx.JSON(http.StatusAccepted, gin.H{"code": rolexerror.CodeOk, "data": data})
	return
}

func HttpErrorResponse(ctx *gin.Context, err error) {
	rerror, ok := err.(*rolexerror.RolexError)
	if !ok {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"code": rolexerror.CodeUndefined, "data": err.Error()})
		return
	}

	var httpCode int
	switch rerror.Code {
	case rolexerror.CodeNetworkPredefined:
		httpCode = http.StatusForbidden
	case rolexerror.CodeListContainerParamError,
		rolexerror.CodePatchContainerParamError,
		rolexerror.CodePatchContainerMethodUndefined,
		rolexerror.CodeDeleteContainerParamError,
		rolexerror.CodeDeleteContainerMethodUndefined,
		rolexerror.CodeListImageParamError,
		rolexerror.CodeConnectNetworkParamError,
		rolexerror.CodeConnectNetworkMethodError,
		rolexerror.CodeCreateNetworkParamError,
		rolexerror.CodeInspectNetworkParamError,
		rolexerror.CodeListNetworkParamError,
		rolexerror.CodeUpdateNodeParamError,
		rolexerror.CodeUpdateServiceParamError,
		rolexerror.CodeCreateServiceParamError,
		rolexerror.CodeScaleServiceParamError,
		rolexerror.CodeListTaskParamError,
		rolexerror.CodeCreateStackParamError,
		rolexerror.CodeInvalidGroupId,
		rolexerror.CodeCreateVolumeParamError,
		rolexerror.CodeContainerNotRunning,
		rolexerror.CodeContainerAlreadyRunning,
		rolexerror.CodeAccountCreateParamError,
		rolexerror.CodeAccountLoginParamError,
		rolexerror.CodeAccountGroupAccountsGroupIdNotValidError,
		rolexerror.CodeAccountAccoutGroupsAccountIdNotValidError,
		rolexerror.CodeAccountCreateGroupParamError,
		rolexerror.CodeAccountUpdateGroupParamError,
		rolexerror.CodeAccountGrantServicePermissionParamError,
		rolexerror.CodeAccountRevokeServicePermissionParamError:
		httpCode = http.StatusBadRequest

	case rolexerror.CodeContainerNotFound,
		rolexerror.CodeNetworkNotFound,
		rolexerror.CodeNetworkOrContainerNotFound,
		rolexerror.CodeAccountGroupAccountsNotFoundError,
		rolexerror.CodeAccountGetGroupGroupIdNotFoundError:
		httpCode = http.StatusNotFound

	default:
		httpCode = http.StatusServiceUnavailable
	}

	ctx.JSON(httpCode, gin.H{"code": rerror.Code, "data": err.Error()})
	return
}
