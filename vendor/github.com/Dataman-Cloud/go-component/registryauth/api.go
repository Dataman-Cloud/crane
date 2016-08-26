package registryauth

import (
	"github.com/Dataman-Cloud/go-component/utils/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
)

type HubApi struct {
	DbClient *gorm.DB
}

var hubApi HubApi

type RegistryAuth struct {
	ID        uint64 `json:"Id"`
	Name      string `json:"Name" gorm:"not null"`
	Username  string `json:"Username" gorm:"not null"`
	Password  string `json:"Password" gorm:"not null"`
	AccountId uint64 `json:"Accountid" gorm:"not null"`
}

func GetHubApi() *HubApi {
	if &hubApi == nil {
		return &HubApi{}
	}
	return &hubApi
}

func (hubApi *HubApi) MigriateRegistryAuth() {
	hubApi.DbClient = db.DB()
	hubApi.DbClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&RegistryAuth{}).AddUniqueIndex("idx_name_userid", "name", "account_id")
}

func (hubApi *HubApi) Create(auth *RegistryAuth) error {
	return hubApi.DbClient.Save(auth).Error
}

func (hubApi *HubApi) List(registryAuth *RegistryAuth) ([]RegistryAuth, error) {
	var registryAuths []RegistryAuth
	err := hubApi.DbClient.Where(registryAuth).Find(&registryAuths).Error
	return registryAuths, err
}

func (hubApi *HubApi) Delete(registryAuth *RegistryAuth) error {
	return hubApi.DbClient.Delete(registryAuth).Error
}
