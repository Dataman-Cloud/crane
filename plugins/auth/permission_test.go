package auth

import (
	"github.com/Dataman-Cloud/rolex/plugins/auth"
	"testing"
)

func TestNewPerm(T *testing.T) {
	p := auth.NewPermission("x")
	if p.Display != "x" {
		T.Error("p.Display should be 'x'")
	}

	if p.Perm != 2 {
		T.Error("p.Value should be 2")
	}
}

func TestNormalize(T *testing.T) {
	p := auth.Permission{Display: "x"}
	p = p.Normalize()
	if p.Perm != 2 {
		T.Error("p.Value should be 2")
	}
}

func TestPermGreaterThanPerm(T *testing.T) {
	p := auth.NewPermission("x")
	greaterPerms := auth.PermGreaterOrEqualThan(p)
	if len(greaterPerms) != 1 {
		T.Error("permissions greater than x should be only x itself")
	}

	p = auth.NewPermission("r")
	greaterPerms = auth.PermGreaterOrEqualThan(p)
	if len(greaterPerms) != 3 {
		T.Error("permissions greater than r should be r, w, x")
	}
}

func TestPermLessThanPerm(T *testing.T) {
	p := auth.NewPermission("x")
	greaterPerms := auth.PermLessOrEqualThan(p)
	if len(greaterPerms) != 3 {
		T.Error("permissions less than x should be r, w, x")
	}

	p = auth.NewPermission("r")
	greaterPerms = auth.PermLessOrEqualThan(p)
	if len(greaterPerms) != 1 {
		T.Error("permissions greater than r should be r itself")
	}
}
