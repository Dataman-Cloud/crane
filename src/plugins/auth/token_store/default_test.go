package token_store

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func NewTestDefaultStore() *Default {
	d := new(Default)
	d.Store = map[string]*tokenStore{}
	return d
}

func TestNewDefaultStore(t *testing.T) {
	if d := NewDefaultStore(); d != nil {
		t.Log("pass")
	} else {
		t.Error("failed")
	}
}

func TestSet(t *testing.T) {
	d := NewTestDefaultStore()
	if err := d.Set(new(gin.Context), "test", "1", time.Now()); err != nil {
		t.Error("failed")
	} else {
		t.Log("pass")
	}
}

func TestGet(t *testing.T) {
	d := NewTestDefaultStore()
	d.Store["test"] = &tokenStore{AccountId: "1", ExpireAt: time.Now().Add(time.Duration(time.Second * 600))}
	if _, err := d.Get(new(gin.Context), "test"); err != nil {
		t.Error(err)
	} else {
		t.Log("pass")
	}
}

func TestDel(t *testing.T) {
	d := NewTestDefaultStore()
	d.Store["test"] = &tokenStore{AccountId: "1"}
	if err := d.Del(new(gin.Context), "test"); err != nil {
		t.Error(err)
	} else {
		t.Log("pass")
	}
}
