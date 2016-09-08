package license

import (
	"github.com/Dataman-Cloud/crane/src/plugins/apiplugin"
	"github.com/Dataman-Cloud/crane/src/utils/db"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
)

func init() {
	licenseApi := &LicenseApi{DbClient: db.DB()}
	licenseApi.MigriateSetting()

	apiPlugin := &apiplugin.ApiPlugin{
		Name:         apiplugin.License,
		Dependencies: []string{apiplugin.Db},
		Instance:     licenseApi,
	}

	apiplugin.Add(apiPlugin)
}

type LicenseApi struct {
	DbClient *gorm.DB
}

func (licenseApi *LicenseApi) ApiRegister(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	licenseV1 := router.Group("/license/v1", middlewares...)
	{
		licenseV1.GET("/license", licenseApi.Get)
		licenseV1.POST("/license", licenseApi.Create)
	}
}
