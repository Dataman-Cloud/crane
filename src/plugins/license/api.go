package license

import (
	"strconv"

	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"
	"github.com/Dataman-Cloud/rolex/src/util/rolexgin"

	"github.com/Dataman-Cloud/go-component/encrypt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

var key = "abcdefghijklmnopqrstuvwx"

func (licenseApi *LicenseApi) MigriateSetting() {
	licenseApi.DbClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Setting{})
}

func (licenseApi *LicenseApi) Create(ctx *gin.Context) {
	var err error

	var setting Setting
	if err = ctx.BindJSON(&setting); err != nil {
		log.Errorf("invalid data error: %v", err)
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeLicenseCreateLicenseError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if _, err = encrypt.Decrypt(key, setting.License); err != nil {
		log.Errorf("invalid license error: %v", err)
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeLicenseCreateLicenseError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	var objSetting Setting
	if err = licenseApi.DbClient.
		Select("license").
		First(&objSetting).
		Error; err != nil {
		log.Errorf("get license error: %v", err)
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeLicenseCreateLicenseError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	objSetting.License = setting.License
	if err = licenseApi.DbClient.
		Model(&Setting{}).
		Select("license").
		Update(&objSetting).
		Error; err != nil {
		log.Errorf("update license error: %v", err)
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeLicenseCreateLicenseError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	rolexgin.HttpOkResponse(ctx, "update license success")
}

func (licenseApi *LicenseApi) Get(ctx *gin.Context) {
	var err error

	var objSetting Setting
	if err = licenseApi.DbClient.
		Select("license").
		First(&objSetting).
		Error; err != nil {
		log.Errorf("get license error: %v", err)
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeLicenseGetLicenseError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if lc, err := encrypt.Decrypt(key, objSetting.License); err != nil {
		log.Errorf("invalid license error: %v", err)
		rolexerr := rolexerror.NewRolexError(rolexerror.CodeLicenseGetLicenseError, err.Error())
		rolexgin.HttpErrorResponse(ctx, rolexerr)
		return
	} else {
		objSetting.License = lc
	}

	rolexgin.HttpOkResponse(ctx, objSetting)
}

func (licenseApi *LicenseApi) GetLicenseValidity() (uint64, error) {
	var err error

	var objSetting Setting
	if err = licenseApi.DbClient.
		Select("license").
		First(&objSetting).
		Error; err != nil {
		log.Errorf("get license error: %v", err)
		return 0, err
	}

	lc, err := encrypt.Decrypt(key, objSetting.License)
	if err != nil {
		return 0, err
	}

	l, err := strconv.ParseUint(lc, 10, 64)
	if err != nil {
		return 0, nil
	}

	return l, nil
}
