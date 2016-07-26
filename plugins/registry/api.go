package registry

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/Dataman-Cloud/rolex/plugins/auth"
	"github.com/Dataman-Cloud/rolex/plugins/auth/authenticators"
	"github.com/Dataman-Cloud/rolex/util/config"
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
