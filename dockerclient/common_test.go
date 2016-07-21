package dockerclient

import (
	"github.com/Dataman-Cloud/rolex/dockerclient"
	"testing"
)

func TestPermissionEqual(T *testing.T) {
	p1 := dockerclient.Permission{Permission: "ro", Group: "1"}
	p2 := dockerclient.Permission{Permission: "ro", Group: "1"}

	if p1.Equal(p2) {
		T.Log("p1 equal p2")
	} else {
		T.Error("p1 should equal p2")
	}
}

func TestPermissiosFromLabel(T *testing.T) {
	label := "group_id-ro"
	permissions := dockerclient.PermissionsFromLabel(label)

	if len(permissions) != 1 {
		T.Error("label group_id-ro should contain 1 permissions")
	}

	if permissions[0].Permission != "ro" {
		T.Error("permission should be read only")
	}

	if permissions[0].Group != "group_id" {
		T.Error("group id should be group_id")
	}

	label = "group_id-ro,group_id1-wo"
	permissions = dockerclient.PermissionsFromLabel(label)

	if len(permissions) != 2 {
		T.Error("label group_id-ro should contain 2 permissions")
	}

	if permissions[1].Permission != "wo" {
		T.Error("permission should be write only")
	}

	if permissions[1].Group != "group_id1" {
		T.Error("group id should be group_id")
	}

	label = "group_id-ro,,group_id1-wo"
	permissions = dockerclient.PermissionsFromLabel(label)
	if len(permissions) != 2 {
		T.Error("label group_id-ro should contain 2 permissions")
	}

	label = "group_id-ro,  ,group_id1-wo"
	permissions = dockerclient.PermissionsFromLabel(label)
	if len(permissions) != 2 {
		T.Error("label group_id-ro should contain 2 permissions")
	}

	label = "group_id-ro, foobar ,group_id1-wo"
	permissions = dockerclient.PermissionsFromLabel(label)
	if len(permissions) != 2 {
		T.Error("label group_id-ro should contain 2 permissions")
	}
}

func TestLabelFromPermission(T *testing.T) {
	permissions := []dockerclient.Permission{
		{Permission: "ro", Group: "1"},
	}

	if dockerclient.PermissionsToLabel(permissions) != "1-ro" {
		T.Error("permissions label should be 1-ro")
	}

	permissions = []dockerclient.Permission{
		{Permission: "ro", Group: "1"},
		{Permission: "rw", Group: "2"},
	}

	if dockerclient.PermissionsToLabel(permissions) != "1-ro,2-rw" {
		T.Error("permissions label should be 2-rw,1-ro")
	}
}

func TestPermissionsInclude(T *testing.T) {
	permissions := []dockerclient.Permission{
		{Permission: "ro", Group: "1"},
	}

	permission := dockerclient.Permission{Permission: "ro", Group: "1"}

	if !dockerclient.PermissionsInclude(permissions, permission) {
		T.Error("permissions should include permssion")
	}
}

func TestPermissionsIndex(T *testing.T) {
	permissions := []dockerclient.Permission{
		{Permission: "ro", Group: "0"},
		{Permission: "ro", Group: "1"},
	}

	permission := dockerclient.Permission{Permission: "ro", Group: "1"}

	if dockerclient.PermissionsIndex(permissions, permission) != 1 {
		T.Error("index of permission in permissions should be 1")
	}
}
