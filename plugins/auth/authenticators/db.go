package authenticators

import (
	"errors"

	"github.com/Dataman-Cloud/rolex/plugins/auth"
	"github.com/Dataman-Cloud/rolex/util/db"

	"github.com/jinzhu/gorm"
)

var err error

type DbAuthenicator struct {
	DbClient *gorm.DB
	auth.Authenticator
}

func NewDBAuthenticator() *DbAuthenicator {
	return &DbAuthenicator{DbClient: db.DB()}
}

func (db *DbAuthenicator) ModificationAllowed() bool {
	return true
}

func (db *DbAuthenicator) Login(a *auth.Account) (string, error) {
	if err = db.DbClient.
		Where("email = ? AND password = ?", a.Email, a.Password).
		First(a).Error; err != nil {
		return "", err
	}

	return auth.GenToken(a), nil
}

func (db *DbAuthenicator) Accounts(filter auth.AccountFilter) (*[]auth.Account, error) {
	var auths []auth.Account
	if err = db.DbClient.Where(&filter).Find(&auths).Error; err != nil {
		return nil, err
	}

	return &auths, nil
}

func (db *DbAuthenicator) Account(idOrEmail interface{}) (*auth.Account, error) {
	var acc auth.Account

	if err = db.DbClient.
		Where("id = ?", idOrEmail).
		Or("email = ?", idOrEmail).
		First(&acc).
		Error; err != nil {
		return nil, err
	}

	return &acc, nil
}

func (db *DbAuthenicator) DeleteAccount(a *auth.Account) error {
	return db.DbClient.Delete(a).Error
}

func (db *DbAuthenicator) UpdaetAccount(a *auth.Account) error {
	return db.DbClient.Save(a).Error
}

func (db *DbAuthenicator) CreateAccount(a *auth.Account) error {
	return db.DbClient.Save(a).Error
}

func (db *DbAuthenicator) Groups(group auth.GroupFilter) (*[]auth.Group, error) {
	return nil, nil
}

func (db *DbAuthenicator) Group(id uint64) (*auth.Group, error) {
	var group auth.Group
	if err = db.DbClient.Where("id = ?", id).Find(&group).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (db *DbAuthenicator) CreateGroup(g *auth.Group) error {
	return db.DbClient.Save(g).Error
}

func (db *DbAuthenicator) DeleteGroup(groupId uint64) error {
	var accounts []auth.AccountGroup
	if err = db.DbClient.Where("group_id = ?", groupId).Find(&accounts).Error; err != nil {
		return err
	}

	if len(accounts) > 0 {
		return errors.New("group contains some account")
	}

	return db.DbClient.Where("id = ?", groupId).Delete(&auth.Group{}).Error
}

func (db *DbAuthenicator) UpdateGroup(g *auth.Group) error {
	return db.DbClient.Model(g).Update(g).Error
}

func (db *DbAuthenicator) JoinGroup(accountId, groupId uint64) error {
	var accounts []auth.AccountGroup

	if err = db.DbClient.
		Where("account_id = ? AND group_id = ?", accountId, groupId).
		Find(&accounts).
		Error; err != nil {
		return err
	}
	if len(accounts) > 0 {
		return errors.New("already exist")
	}

	if err = db.DbClient.Save(&auth.AccountGroup{AccountId: accountId, GroupId: groupId}).Error; err != nil {
		return err
	}

	return nil
}

func (db *DbAuthenicator) LeaveGroup(accountId, groupId uint64) error {
	var accounts []auth.AccountGroup

	if err = db.DbClient.
		Where("account_id = ? AND group_id = ?", accountId, groupId).
		Find(&accounts).
		Error; err != nil {
		return err
	}
	if len(accounts) == 0 {
		return errors.New("account group non-existent")
	}

	if err = db.DbClient.
		Where("account_id = ? AND group_id = ?", accountId, groupId).
		Delete(&auth.AccountGroup{}).
		Error; err != nil {
		return err
	}

	return nil
}
