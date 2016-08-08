package license

import (
	"github.com/Dataman-Cloud/rolex/src/util/db"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
)

type LicenseApi struct {
	DbClient *gorm.DB
}

func (licenseApi *LicenseApi) RegisterApiForLicense(router *gin.Engine,
	middlewares ...gin.HandlerFunc) {
	licenseApi.DbClient = db.DB()
	licenseApi.MigriateSetting()

	licenseV1 := router.Group("/license/v1", middlewares...)
	{
		licenseV1.GET("/license", licenseApi.Get)
		licenseV1.POST("/license", licenseApi.Create)
	}
}
