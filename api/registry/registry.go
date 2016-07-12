package registry

import (
	"fmt"
	"net/http"

	"github.com/Dataman-Cloud/rolex/util/config"
	"github.com/gin-gonic/gin"
)

type Registry struct {
	Config *config.Config
}

func (registry *Registry) Token(ctx *gin.Context) {
	username, password, _ := ctx.Request.BasicAuth()
	authenticated := authenticate(username, password)

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
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": rawToken})
}

// TODO check if account valid here
func authenticate(principal, password string) bool {
	return true
}
