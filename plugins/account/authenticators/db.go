package authenticators

import (
	"github.com/Dataman-Cloud/rolex/plugins/account"
	"github.com/Dataman-Cloud/rolex/util/db"

	"github.com/jinzhu/gorm"
)

type DbAuthenicator struct {
	DbClient *gorm.DB
	account.Authenticator
}

func NewDBAuthenticator() *DbAuthenicator {
	return &DbAuthenicator{DbClient: db.DB()}
}

func (db *DbAuthenicator) ModificationAllowed() bool {
	return true
}

func (db *DbAuthenicator) Login(a *account.Account) (string, error) {
	if err := db.DbClient.
		Where("email = ? AND password = ?", a.Email, a.Password).
		First(a).Error; err != nil {
		return "", err
	}

	return account.GenToken(a), nil
}

func (db *DbAuthenicator) Accounts(filter account.AccountFilter) (*[]account.Account, error) {
	var accounts []account.Account
	if err := db.DbClient.Where(&filter).Find(&accounts).Error; err != nil {
		return nil, err
	}

	return &accounts, nil
}

func (db *DbAuthenicator) Account(idOrEmail interface{}) (*account.Account, error) {
	var acc account.Account

	if err := db.DbClient.
		Where("id = ?", idOrEmail).
		Or("email = ?", idOrEmail).
		First(&acc).
		Error; err != nil {
		return nil, err
	}

	return &acc, nil
}

func (db *DbAuthenicator) DeleteAccount(a *account.Account) error {
	return db.DbClient.Delete(a).Error
}

func (db *DbAuthenicator) UpdaetAccount(a *account.Account) error {
	return db.DbClient.Save(a).Error
}

func (db *DbAuthenicator) CreateAccount(a *account.Account) error {
	return db.DbClient.Save(a).Error
}

func (db *DbAuthenicator) Groups(group account.GroupFilter) (*[]account.Group, error) {
	return nil, nil
}

func (db *DbAuthenicator) Group(id uint64) (*account.Group, error) {
	var group account.Group
	if err := db.DbClient.Where("id = ?", id).Find(&group).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (db *DbAuthenicator) CreateGroup(g *account.Group) error {
	return db.DbClient.Save(g).Error
}

func (db *DbAuthenicator) DeleteGroup(groupId uint64) error {
	return db.DbClient.Where("id = ?", groupId).Delete(&account.Group{}).Error
}

func (db *DbAuthenicator) UpdateGroup(g *account.Group) error {
	return db.DbClient.Model(g).Update(g).Error
}

func (db *DbAuthenicator) JoinGroup(g *account.Group, a *account.Account) error {
	return nil
}

func (db *DbAuthenicator) LeaveGroup(g *account.Group, a *account.Account) error {
	return nil
}
