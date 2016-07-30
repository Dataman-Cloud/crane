package registry

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/Dataman-Cloud/rolex/src/plugins/auth"
	"github.com/Dataman-Cloud/rolex/src/plugins/auth/authenticators"
	"github.com/Dataman-Cloud/rolex/src/util/config"

	"github.com/docker/distribution/manifest/schema1"
	"github.com/gin-gonic/gin"
)

const manifestPattern = `^application/vnd.docker.distribution.manifest.v\d`

type Registry struct {
	Config *config.Config
}

func (registry *Registry) Token(ctx *gin.Context) {
	username, password, _ := ctx.Request.BasicAuth()
	authenticated := registry.Authenticate(username, password)

	service := ctx.Query("service")
	scope := ctx.Query("scope")

	if len(scope) == 0 && !authenticated {
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	accesses := ParseResourceActions(scope)
	for _, access := range accesses {
		FilterAccess(username, authenticated, access)
	}

	//create token
	rawToken, err := MakeToken(registry.Config, username, service, accesses)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": rawToken})
}

func (registry *Registry) Authenticate(principal, password string) bool {
	var authenticator auth.Authenticator
	if registry.Config.AccountAuthenticator == "db" {
	} else if registry.Config.AccountAuthenticator == "ldap" {
	} else {
		authenticator = authenticators.NewDefaultAuthenticator()
	}

	_, err := authenticator.Login(&auth.Account{Email: principal, Password: password})
	if err != nil {
		return false
	}
	return true
}

func (registry *Registry) Notifications(ctx *gin.Context) {
	notification := &Notification{}
	if err := ctx.BindJSON(&notification); err != nil {
		switch jsonErr := err.(type) {
		case *json.SyntaxError:
			fmt.Printf("Notification JSON syntax error at byte %v: %s", jsonErr.Offset, jsonErr.Error())
		case *json.UnmarshalTypeError:
			fmt.Printf("Unexpected type at by type %v. Expected %s but received %s.",
				jsonErr.Offset, jsonErr.Type, jsonErr.Value)
		}
	}

	for _, e := range notification.Events {
		matched, _ := regexp.MatchString(manifestPattern, e.Target.MediaType)
		if matched && strings.HasPrefix(ctx.Request.UserAgent(), "docker") {
			fmt.Println(e)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (registry *Registry) TagList(ctx *gin.Context) {
	account_, found := ctx.Get("account")
	if !found {
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	account := account_.(auth.Account)

	resp, err := registry.RegistryAPIGet(fmt.Sprintf("%s/%s/tags/list", ctx.Param("namespace"), ctx.Param("image")), account.Email)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	var respBody struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": respBody})
}

func (registry *Registry) GetManifests(ctx *gin.Context) {
	account_, found := ctx.Get("account")
	if !found {
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	account := account_.(auth.Account)

	resp, err := registry.RegistryAPIGet(fmt.Sprintf("%s/%s/manifests/%s", ctx.Param("namespace"), ctx.Param("image"), ctx.Param("reference")), account.Email)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	var manifest schema1.Manifest
	err = json.Unmarshal(resp, &manifest)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": manifest})
}

func (registry *Registry) UpdateManifests(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (registry *Registry) DeleteManifests(ctx *gin.Context) {
	account_, found := ctx.Get("account")
	if !found {
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	account := account_.(auth.Account)

	_, err := registry.RegistryAPIDelete(fmt.Sprintf("%s/%s/manifests/%s", ctx.Param("namespace"), ctx.Param("image"), ctx.Param("reference")), account.Email)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "success"})
}

func (registry *Registry) Catalog(ctx *gin.Context) {
	account_, found := ctx.Get("account")
	if !found {
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	account := account_.(auth.Account)

	resp, err := registry.RegistryAPIGet("_catalog", account.Email)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	var respBody struct {
		Repositories []string `json:"repositories"`
	}
	err = json.Unmarshal(resp, &respBody)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": respBody})
}
