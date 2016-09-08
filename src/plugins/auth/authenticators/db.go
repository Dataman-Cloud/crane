package authenticators

import (
	"crypto/md5"
	"errors"
	"fmt"
	"strconv"

	"github.com/Dataman-Cloud/crane/src/plugins/auth"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/db"
	"github.com/Dataman-Cloud/crane/src/utils/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
)

var err error

type DbAuthenicator struct {
	DbClient *gorm.DB
	auth.Authenticator
}

func NewDBAuthenticator() *DbAuthenicator {
	authenticator := &DbAuthenicator{DbClient: db.DB()}
	authenticator.MigriateTable()
	return authenticator
}

func (authenticator *DbAuthenicator) MigriateTable() {
	authenticator.DbClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&auth.Account{})
	authenticator.DbClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&auth.Group{})
	authenticator.DbClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&auth.AccountGroup{})
}

func (db *DbAuthenicator) ModificationAllowed() bool {
	return true
}

func (db *DbAuthenicator) Login(a *auth.Account) (string, error) {
	var account auth.Account
	if err = db.DbClient.
		Select("id, title, email, phone, login_at, password").
		Where("email = ?", a.Email).
		First(&account).Error; err != nil {
		return "", cranerror.NewError(auth.CodeAccountLoginFailedEmailNotValidError, err.Error())
	} else if account.Password != db.EncryptPassword(a.Password) {
		return "", cranerror.NewError(auth.CodeAccountLoginFailedPasswordNotValidError, "Invalid Password")
	}

	a.ID = account.ID
	return auth.GenToken(&account), nil
}

func (db *DbAuthenicator) Accounts(listOptions model.ListOptions) (*[]auth.Account, error) {
	var auths []auth.Account
	if err = db.DbClient.
		Select("id, title, email, phone, login_at").
		Offset(listOptions.Offset).
		Limit(listOptions.Limit).
		Find(&auths).Error; err != nil {
		return nil, err
	}

	return &auths, nil
}

func (db *DbAuthenicator) Account(idOrEmail interface{}) (*auth.Account, error) {
	var acc auth.Account

	if id, err := strconv.ParseUint(fmt.Sprintf("%v", idOrEmail), 10, 64); err == nil {
		if err = db.DbClient.
			Select("id, title, email, phone, login_at").
			Where("id = ?", id).
			First(&acc).
			Error; err != nil {
			return nil, err
		}
	} else {
		if err = db.DbClient.
			Select("id, title, email, phone, login_at").
			Where("email = ?", idOrEmail).
			First(&acc).
			Error; err != nil {
			return nil, err
		}
	}

	return &acc, nil
}

func (db *DbAuthenicator) DeleteAccount(a *auth.Account) error {
	return db.DbClient.Delete(a).Error
}

func (db *DbAuthenicator) UpdaetAccount(a *auth.Account) error {
	return db.DbClient.Save(a).Error
}

func (db *DbAuthenicator) CreateAccount(groupId uint64, a *auth.Account) error {
	tx := db.DbClient.Begin()
	gcount := 0
	if err = tx.Model(&auth.Group{}).Where("id = ?", groupId).Count(&gcount).Error; err != nil {
		return err
	}

	if gcount == 0 {
		return errors.New("not found group")
	}

	acount := 0
	if err = tx.Model(&auth.Account{}).Where("email = ?", a.Email).Count(&acount).Error; err != nil {
		return err
	}

	if acount > 0 {
		return errors.New("email already exists")
	}

	if err = tx.Save(a).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Save(&auth.AccountGroup{
		AccountId: a.ID,
		GroupId:   groupId,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (db *DbAuthenicator) Groups(listOptions model.ListOptions) (*[]auth.Group, error) {
	var groups []auth.Group
	if err = db.DbClient.
		Offset(listOptions.Offset).
		Limit(listOptions.Limit).
		Where(listOptions.Filter).
		Find(&groups).Error; err != nil {
		return nil, err
	}
	return &groups, nil
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

func (db *DbAuthenicator) GroupAccounts(listOptions model.ListOptions) (*[]auth.Account, error) {
	var accounts []auth.Account
	var accountGroup []auth.AccountGroup
	if err := db.DbClient.
		Where(listOptions.Filter).
		Offset(listOptions.Offset).
		Limit(listOptions.Limit).
		Find(&accountGroup).
		Error; err != nil {
		return nil, err
	}

	for _, ag := range accountGroup {
		var account auth.Account
		if err := db.DbClient.
			Select("id, title, email, phone, login_at").
			Where("id = ?", ag.AccountId).
			Find(&account).
			Error; err == nil {
			accounts = append(accounts, account)
		}
	}
	return &accounts, nil
}

func (db *DbAuthenicator) AccountGroups(listOptions model.ListOptions) (*[]auth.Group, error) {
	var groups []auth.Group
	var accountGroup []auth.AccountGroup
	if err := db.DbClient.
		Where(listOptions.Filter).
		Offset(listOptions.Offset).
		Limit(listOptions.Limit).
		Find(&accountGroup).
		Error; err != nil {
		return nil, err
	}

	for _, ag := range accountGroup {
		var group auth.Group
		if err := db.DbClient.Where("id = ?", ag.GroupId).Find(&group).Error; err == nil {
			groups = append(groups, group)
		}
	}
	return &groups, nil
}

func (db *DbAuthenicator) EncryptPassword(password string) string {
	pw := fmt.Sprintf("dataman-crane%xdataman-crane", md5.Sum([]byte(password)))
	return fmt.Sprintf("%x", md5.Sum([]byte(pw)))
}
