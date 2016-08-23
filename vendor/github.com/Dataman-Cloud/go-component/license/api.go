package license

import (
	"strconv"

	"github.com/Dataman-Cloud/go-component/utils/dmerror"
	"github.com/Dataman-Cloud/go-component/utils/dmgin"

	"github.com/Dataman-Cloud/go-component/utils/encrypt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

const (
	//license
	CodeLicenseGetLicenseError    = "503-16001"
	CodeLicenseCreateLicenseError = "503-16002"
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
		rolexerr := dmerror.NewError(CodeLicenseCreateLicenseError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if sk, err := encrypt.Decrypt(key, setting.License); err != nil {
		log.Errorf("invalid license error: %v", err)
		rolexerr := dmerror.NewError(CodeLicenseCreateLicenseError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	} else {
		if _, err = strconv.ParseUint(sk, 10, 64); err != nil {
			log.Errorf("invalid license error: %v", err)
			rolexerr := dmerror.NewError(CodeLicenseCreateLicenseError, err.Error())
			dmgin.HttpErrorResponse(ctx, rolexerr)
			return
		}
	}

	var objSetting Setting
	if err = licenseApi.DbClient.
		Select("license").
		First(&objSetting).
		Error; err != nil {
		objSetting.License = setting.License
		if err = licenseApi.DbClient.
			Model(&Setting{}).
			Select("license").
			Save(&objSetting).
			Error; err != nil {
			log.Errorf("update license error: %v", err)
			rolexerr := dmerror.NewError(CodeLicenseCreateLicenseError, err.Error())
			dmgin.HttpErrorResponse(ctx, rolexerr)
			return
		}
		dmgin.HttpOkResponse(ctx, "update license success")
		return
	}

	objSetting.License = setting.License
	if err = licenseApi.DbClient.
		Model(&Setting{}).
		Select("license").
		Update(&objSetting).
		Error; err != nil {
		log.Errorf("update license error: %v", err)
		rolexerr := dmerror.NewError(CodeLicenseCreateLicenseError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	dmgin.HttpOkResponse(ctx, "update license success")
}

func (licenseApi *LicenseApi) Get(ctx *gin.Context) {
	var err error

	var objSetting Setting
	if err = licenseApi.DbClient.
		Select("license").
		First(&objSetting).
		Error; err != nil {
		log.Errorf("get license error: %v", err)
		rolexerr := dmerror.NewError(CodeLicenseGetLicenseError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	}

	if lc, err := encrypt.Decrypt(key, objSetting.License); err != nil {
		log.Errorf("invalid license error: %v", err)
		rolexerr := dmerror.NewError(CodeLicenseGetLicenseError, err.Error())
		dmgin.HttpErrorResponse(ctx, rolexerr)
		return
	} else {
		objSetting.License = lc
	}

	dmgin.HttpOkResponse(ctx, objSetting)
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
