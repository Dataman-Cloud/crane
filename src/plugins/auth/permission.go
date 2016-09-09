package auth

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PERMISSION_LABEL_PREFIX = "crane.reserved.permissions"
)

type Permission struct {
	Perm    int    `json:"Perm"`
	Display string `json:"Display"`
}

type GroupPermission struct {
	Permission Permission `json:"Permission"`
	GroupID    uint64     `json:"GroupID"`
}

var (
	PermReadOnly  Permission = Permission{Perm: 0, Display: "r"}
	PermReadWrite Permission = Permission{Perm: 1, Display: "w"}
	PermAdmin     Permission = Permission{Perm: 2, Display: "x"}

	Perms []Permission = []Permission{PermReadOnly, PermReadWrite, PermAdmin}
)

func (p Permission) Normalize() Permission {
	return NewPermission(p.Display)
}

func NewPermission(display string) Permission {
	for _, v := range Perms {
		if v.Display == display {
			return v
		}
	}
	return Permission{}
}

func PermGreaterOrEqualThan(p Permission) []Permission {
	p = p.Normalize()
	var perms []Permission
	for _, v := range Perms {
		if v.Perm >= p.Perm {
			perms = append(perms, v)
		}
	}

	return perms
}

func PermLessOrEqualThan(p Permission) []Permission {
	p = p.Normalize()
	var perms []Permission
	for _, v := range Perms {
		if v.Perm <= p.Perm {
			perms = append(perms, v)
		}
	}

	return perms
}

func PermissionRevokeLabelKeysFromPermissionId(permissionId string) []string {
	var param struct {
		GroupID uint64 `json:"GroupID"`
		Perm    string `json:"Perm"`
	}
	param.GroupID, _ = strconv.ParseUint(strings.SplitN(permissionId, "-", 2)[0], 10, 64)
	param.Perm = strings.SplitN(permissionId, "-", 2)[1]

	labels := make([]string, 0)
	gp := GroupPermission{GroupID: param.GroupID, Permission: Permission{Display: param.Perm}}
	for _, perm := range PermGreaterOrEqualThan(gp.Permission) {
		labels = append(labels, fmt.Sprintf("%s.%d.%s", PERMISSION_LABEL_PREFIX, gp.GroupID, perm.Display))
	}

	return labels
}

func PermissionGrantLabelsPairFromGroupIdAndPerm(groupId uint64, perm string) map[string]string {
	labels := make(map[string]string, 0)
	gp := GroupPermission{GroupID: groupId, Permission: Permission{Display: perm}}
	for _, perm := range PermLessOrEqualThan(gp.Permission) {
		labels[fmt.Sprintf("%s.%d.%s", PERMISSION_LABEL_PREFIX, gp.GroupID, perm.Display)] = "true"
	}

	return labels
}
