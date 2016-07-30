package auth

import (
	"github.com/Dataman-Cloud/rolex/src/plugins/auth"
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

func TestPermissionRevokeLabelsFromPermissionId(T *testing.T) {
	labels := auth.PermissionRevokeLabelKeysFromPermissionId("gid-r")

	if len(labels) != 3 {
		T.Error("labels from permission gid-x should have 3 items")
	}

	if labels[0] == "com.rolex.permissions.gid.r" {
		T.Error("labels from permission 1 item is r")
	}

	if labels[1] == "com.rolex.permissions.gid.w" {
		T.Error("labels from permission 2 item is w")
	}

	if labels[2] == "com.rolex.permissions.gid.x" {
		T.Error("labels from permission 3 item is x")
	}
}

func TestPermissionGrantLablePairsFromGroupIdAndPerm(T *testing.T) {
	labels := auth.PermissionGrantLabelsPairFromGroupIdAndPerm(1, "x")
	if labels["com.rolex.permissions.1.x"] != "true" {
		T.Error("group 1 should have permission x")
	}

	labels = auth.PermissionGrantLabelsPairFromGroupIdAndPerm(1, "w")
	if labels["com.rolex.permissions.1.r"] != "true" {
		T.Error("group 1 should have permission r")
	}

	if labels["com.rolex.permissions.1.w"] != "true" {
		T.Error("group 1 should have permission x")
	}
}
