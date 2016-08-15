package authenticators

import (
	"database/sql/driver"
	"testing"

	"github.com/Dataman-Cloud/rolex/src/model"
	"github.com/Dataman-Cloud/rolex/src/plugins/auth"

	"github.com/erikstmartin/go-testdb"
	"github.com/jinzhu/gorm"
)

type testResult struct {
	lastId       int64
	affectedRows int64
}

func (r testResult) LastInsertId() (int64, error) {
	return r.lastId, nil
}

func (r testResult) RowsAffected() (int64, error) {
	return r.affectedRows, nil
}

func TestLogin(t *testing.T) {
	DB, _ := gorm.Open("testdb", "")
	authenticator := &DbAuthenicator{
		DbClient: DB,
	}

	testdb.SetQueryFunc(func(query string) (driver.Rows, error) {
		columns := []string{"id", "email", "password"}
		result := `
		1,Tim,20
		2,Joe,25
		3,Bob,30`
		return testdb.RowsFromCSVString(columns, result), nil
	})

	var account auth.Account
	if _, err := authenticator.Login(&account); err == nil {
		t.Log("pass")
	} else {
		t.Error(err)
	}
}

func TestAccounts(t *testing.T) {
	DB, _ := gorm.Open("testdb", "")
	authenticator := &DbAuthenicator{
		DbClient: DB,
	}

	testdb.SetQueryFunc(func(query string) (driver.Rows, error) {
		columns := []string{"id", "email", "password"}
		result := `
		1,Tim,20
		2,Joe,25
		3,Bob,30`
		return testdb.RowsFromCSVString(columns, result), nil
	})

	if _, err := authenticator.Accounts(model.ListOptions{}); err == nil {
		t.Log("pass")
	} else {
		t.Error(err)
	}
}

func TestAccount(t *testing.T) {
	DB, _ := gorm.Open("testdb", "")
	authenticator := &DbAuthenicator{
		DbClient: DB,
	}

	testdb.StubQuery(
		`select * from accounts`,
		testdb.RowsFromCSVString([]string{"id"}, `1`),
	)

	if _, err := authenticator.Account(1); err == nil {
		t.Log("pass")
	} else {
		t.Error(err)
	}
}

func TestDeleteAccount(t *testing.T) {
	DB, _ := gorm.Open("testdb", "")
	authenticator := &DbAuthenicator{
		DbClient: DB,
	}

	testdb.SetExecWithArgsFunc(func(query string, args []driver.Value) (result driver.Result, err error) {
		return testResult{1, 1}, nil
	})

	if err := authenticator.DeleteAccount(&auth.Account{}); err == nil {
		t.Log("pass")
	} else {
		t.Error(err)
	}
}

func TestUpdaetAccount(t *testing.T) {
	DB, _ := gorm.Open("testdb", "")
	authenticator := &DbAuthenicator{
		DbClient: DB,
	}

	testdb.SetExecWithArgsFunc(func(query string, args []driver.Value) (result driver.Result, err error) {
		return testResult{1, 1}, nil
	})

	if err := authenticator.UpdaetAccount(&auth.Account{Email: "test", Password: "test"}); err == nil {
		t.Log("pass")
	} else {
		t.Error(err)
	}
}

func TestGroups(t *testing.T) {
	DB, _ := gorm.Open("testdb", "")
	authenticator := &DbAuthenicator{
		DbClient: DB,
	}

	testdb.SetQueryFunc(func(query string) (driver.Rows, error) {
		columns := []string{"id", "name"}
		result := `
		1,Tim
		2,Joe
		3,Bob`
		return testdb.RowsFromCSVString(columns, result), nil
	})

	if _, err := authenticator.Groups(model.ListOptions{}); err == nil {
		t.Log("pass")
	} else {
		t.Error(err)
	}
}

func TestGroup(t *testing.T) {
	DB, _ := gorm.Open("testdb", "")
	authenticator := &DbAuthenicator{
		DbClient: DB,
	}

	testdb.SetQueryFunc(func(query string) (driver.Rows, error) {
		columns := []string{"id", "name"}
		result := `
		1,Tim
		2,Joe
		3,Bob`
		return testdb.RowsFromCSVString(columns, result), nil
	})

	if _, err := authenticator.Group(uint64(1)); err == nil {
		t.Log("pass")
	} else {
		t.Error(err)
	}
}

func TestCreateGroup(t *testing.T) {
	DB, _ := gorm.Open("testdb", "")
	authenticator := &DbAuthenicator{
		DbClient: DB,
	}

	if err := authenticator.CreateGroup(&auth.Group{}); err == nil {
		t.Log("pass")
	} else {
		t.Error(err)
	}
}

func TestUpdateGroup(t *testing.T) {
	DB, _ := gorm.Open("testdb", "")
	authenticator := &DbAuthenicator{
		DbClient: DB,
	}

	if err := authenticator.UpdateGroup(&auth.Group{}); err == nil {
		t.Log("pass")
	} else {
		t.Error(err)
	}
}

func EncryptPassword(t *testing.T) {
	authenticator := &DbAuthenicator{}
	if p := authenticator.EncryptPassword("rolex"); p != "" {
		t.Log("pass")
	} else {
		t.Error(err)
	}
}
