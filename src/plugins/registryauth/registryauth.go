package registryauth

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
)

type RegistryAuth struct {
	ID        uint64 `json:"Id"`
	Name      string `json:"Name" gorm:"not null"`
	Username  string `json:"Username" gorm:"not null"`
	Password  string `json:"Password" gorm:"not null"`
	AccountId uint64 `json:"Accountid" gorm:"not null"`
}

var DbClient *gorm.DB

func Init(dbClient *gorm.DB) {
	DbClient = dbClient
	DbClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&RegistryAuth{}).AddUniqueIndex("idx_name_userid", "name", "account_id")
}

func Create(auth *RegistryAuth) error {
	return DbClient.Save(auth).Error
}

func List(registryAuth *RegistryAuth) ([]RegistryAuth, error) {
	var registryAuths []RegistryAuth
	err := DbClient.Select("id, name, username, account_id").Where(registryAuth).Find(&registryAuths).Error
	return registryAuths, err
}

func Delete(registryAuth *RegistryAuth) error {
	return DbClient.Where("name = ? AND account_id = ?", registryAuth.Name, registryAuth.AccountId).
		Delete(registryAuth).
		Error
}

func Get(name string) (*RegistryAuth, error) {
	var registryAuth RegistryAuth
	err := DbClient.Where("name = ?", name).First(&registryAuth).Error
	return &registryAuth, err
}
