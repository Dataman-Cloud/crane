package api

import (
	"github.com/Dataman-Cloud/crane/src/plugins/auth"
	rauth "github.com/Dataman-Cloud/crane/src/plugins/registryauth"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/httpresponse"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

const (
	CodeCreateRegistryAuthParamError = "400-17001"
	CodeGetRegistryAuthExistError    = "400-17002"
	CodeDeleteRegistryAuthParamError = "400-17003"
	CodeRegistryAuthInvalidUserError = "401-17004"
)

func (api *RegistryAuthApi) Create(ctx *gin.Context) {
	account, ok := ctx.Get("account")
	if !ok {
		log.Error("get registryAuths invalid user")
		httpresponse.Error(ctx, cranerror.NewError(CodeRegistryAuthInvalidUserError, "invalid user"))
		return
	}

	var registryAuth rauth.RegistryAuth
	if err := ctx.BindJSON(&registryAuth); err != nil {
		log.Errorf("create registryAuth param error: %v", err)
		httpresponse.Error(ctx, cranerror.NewError(CodeCreateRegistryAuthParamError, err.Error()))
		return
	}

	rs, err := rauth.List(&rauth.RegistryAuth{Name: registryAuth.Name, AccountId: registryAuth.AccountId})
	if err != nil {
		httpresponse.Error(ctx, err)
		return
	}

	if len(rs) > 0 {
		httpresponse.Error(ctx, cranerror.NewError(CodeGetRegistryAuthExistError, "registryAuth exists"))
		return
	}

	registryAuth.AccountId = account.(auth.Account).ID
	if err := rauth.Create(&registryAuth); err != nil {
		log.Errorf("create registryAuth operation error: %v", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, "create success")
}

func (api *RegistryAuthApi) List(ctx *gin.Context) {
	account, ok := ctx.Get("account")
	if !ok {
		log.Error("get registryAuths invalid user")
		httpresponse.Error(ctx, cranerror.NewError(CodeRegistryAuthInvalidUserError, "invalid user"))
		return
	}

	registryAuth, err := rauth.List(&rauth.RegistryAuth{AccountId: account.(auth.Account).ID})
	if err != nil {
		log.Errorf("get registryAuth by name error: %v", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, registryAuth)
}

func (api *RegistryAuthApi) Delete(ctx *gin.Context) {
	account, ok := ctx.Get("account")
	if !ok {
		log.Error("delete registryAuth invalid user")
		httpresponse.Error(ctx, cranerror.NewError(CodeRegistryAuthInvalidUserError, "invalid user"))
		return
	}

	name := ctx.Param("rauth_name")
	if name == "" {
		log.Errorf("get registryAuth name invalid")
		httpresponse.Error(ctx, cranerror.NewError(CodeDeleteRegistryAuthParamError, "registryAuth name invalid"))
		return
	}

	if err := rauth.Delete(&rauth.RegistryAuth{Name: name, AccountId: account.(auth.Account).ID}); err != nil {
		log.Errorf("delete registryAuth error: %v", err)
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, "delete success")
}
