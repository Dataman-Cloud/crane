package authenticators

import (
	"testing"
)

func TestAccountGet(t *testing.T) {
	d := &Default{}
	a, _ := d.Account(uint64(1))
	if a == nil {
		t.Error("find account with 1 failed")
	}

	a, _ = d.Account("admin@admin.com")
	if a == nil {
		t.Error("find account with admin@admin.com failed")
	}
}
