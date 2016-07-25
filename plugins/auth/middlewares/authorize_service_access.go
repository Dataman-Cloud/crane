package middlewares

import (
	"fmt"
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
			groups, err := a.Authenticator.AccountGroups(&account)
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
				if !ListServiceAuthorization(ctx, groups, permissionRequired) {
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

// add label filter only
func ListServiceAuthorization(ctx *gin.Context, groups *[]auth.Group, permissionRequired auth.Permission) bool {
	labelFilters := make(map[string]string, 0)
	//for _, group := range *groups {
	//labelFilters[fmt.Sprintf("%s.%d.%s", auth.PERMISSION_LABEL_PREFIX, group.ID, permissionRequired.Display)] = "true"
	//}

	// support only one group now as docker label filter doesn't support `OR` operation or `REGEXP` operation
	if len(*groups) > 0 {
		group := (*groups)[0]
		labelFilters[fmt.Sprintf("%s.%d.%s", auth.PERMISSION_LABEL_PREFIX, group.ID, permissionRequired.Display)] = "true"
	}

	ctx.Set("labelFilters", labelFilters)
	return true
}
