package middlewares

import (
	"fmt"
	"github.com/Dataman-Cloud/rolex/model"
	"github.com/Dataman-Cloud/rolex/plugins/auth"

	"github.com/docker/engine-api/types/swarm"
	"github.com/gin-gonic/gin"
)

func AuthorizeServiceAccess(a *auth.AccountApi) func(permissionRequired auth.Permission) gin.HandlerFunc {
	return func(permissionRequired auth.Permission) gin.HandlerFunc {
		permissionRequired = permissionRequired.Normalize()

		return func(ctx *gin.Context) {
			account_, exist := ctx.Get("account")
			if !exist {
				ctx.Abort()
				return
			}

			account := account_.(auth.Account)
			listOptions := model.ListOptions{
				Filter: map[string]interface{}{
					"account_id": account.ID,
				},
			}
			groups, err := a.Authenticator.AccountGroups(listOptions)
			if err != nil {
				ctx.Abort()
				return
			}

			if len(ctx.Param("service_id")) > 0 { // for single service request
				service, err := a.RolexDockerClient.InspectServiceWithRaw(ctx.Param("service_id"))
				if err != nil {
					ctx.Abort()
					return
				}

				if !SingleServiceAuthorizationCheck(service, groups, permissionRequired) {
					ctx.Abort()
					return
				}
			} else { // for list services request
				if !ListServiceAuthorization() {
					ctx.Abort()
					return
				}
			}

			ctx.Next()
		}
	}
}

func SingleServiceAuthorizationCheck(service swarm.Service, groups *[]auth.Group, permissionRequired auth.Permission) bool {
	for _, group := range *groups {
		if service.Spec.Labels[fmt.Sprintf("%s.%d.%s", auth.PERMISSION_LABEL_PREFIX, group.ID, permissionRequired.Display)] == "true" {
			return true
		}
	}
	return false

}

func ListServiceAuthorization() bool {
	return true
}
