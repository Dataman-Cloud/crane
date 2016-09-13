package middlewares

import (
	"github.com/Dataman-Cloud/crane/src/plugins/auth"

	//"github.com/docker/engine-api/types/swarm"
	"github.com/gin-gonic/gin"
)

func AuthorizeServiceAccess() func(permissionRequired auth.Permission) gin.HandlerFunc {
	return func(permissionRequired auth.Permission) gin.HandlerFunc {
		permissionRequired = permissionRequired.Normalize()

		return func(ctx *gin.Context) {
			account_, exist := ctx.Get("account")
			if !exist {
				ctx.Abort()
				return
			}

			account := account_.(auth.Account)
			_ = account
			/*listOptions := model.ListOptions{
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
				service, err := a.CraneDockerClient.InspectServiceWithRaw(ctx.Param("service_id"))
				if err != nil {
					ctx.Abort()
					return
				}

				if !SingleServiceAuthorizationCheck(service, groups, permissionRequired) {
					ctx.Abort()
					return
				}
			} else { // for list services request
				authpass := false
				groupId, err := a.CraneDockerClient.GetStackGroup(ctx.Param("namespace"))
				if err != nil {
					ctx.Abort()
					return
				}
				for _, group := range *groups {
					if group.ID == groupId {
						authpass = true
					}
				}

				if !authpass {
					ctx.Abort()
					return
				}

				if !ListServiceAuthorization(ctx, groupId, permissionRequired) {
					ctx.Abort()
					return
				}

			}*/

			ctx.Next()
		}
	}
}

//func SingleServiceAuthorizationCheck(service swarm.Service, groups *[]auth.Group, permissionRequired auth.Permission) bool {
//for _, group := range *groups {
//if service.Spec.Labels[fmt.Sprintf("%s.%d.%s", auth.PERMISSION_LABEL_PREFIX, group.ID, permissionRequired.Display)] == "true" {
//return true
//}
//}
//return false

//}

//// add label filter only
//func ListServiceAuthorization(ctx *gin.Context, groupId uint64, permissionRequired auth.Permission) bool {
//labelFilters := make(map[string]string, 0)
////for _, group := range *groups {
////labelFilters[fmt.Sprintf("%s.%d.%s", auth.PERMISSION_LABEL_PREFIX, group.ID, permissionRequired.Display)] = "true"
////}

//// support only one group now as docker label filter doesn't support `OR` operation or `REGEXP` operation
//labelFilters[fmt.Sprintf("%s.%d.%s", auth.PERMISSION_LABEL_PREFIX, groupId, permissionRequired.Display)] = "true"

//ctx.Set("labelFilters", labelFilters)
//return true
//}
