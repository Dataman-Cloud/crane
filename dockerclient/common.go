package dockerclient

import (
	"strings"
)

// permissions on label
type Permission struct {
	Permission string `json:"Permission"`
	Group      string `json:"Group"`
}

func (p Permission) Equal(o Permission) bool {
	return p.Permission == o.Permission && o.Group == p.Group
}

func PermissionsFromLabel(label string) []Permission {
	permissions := make([]Permission, 0)
	for _, str := range strings.SplitN(label, ",", -1) {
		pair := strings.SplitN(str, "-", 2)
		if len(pair) != 2 {
			continue
		}
		permissions = append(permissions, Permission{Permission: pair[1], Group: pair[0]})
	}

	return permissions
}

func PermissionsToLabel(permissions []Permission) string {
	pairs := make([]string, 0)
	for _, v := range permissions {
		pairs = append(pairs, strings.Join([]string{v.Group, v.Permission}, "-"))
	}

	return strings.Join(pairs, ",")
}

func PermissionsInclude(permissions []Permission, p Permission) bool {
	for _, v := range permissions {
		if v.Equal(p) {
			return true
		}
	}

	return false
}

func PermissionsIndex(permissions []Permission, p Permission) int {
	for i, v := range permissions {
		if v.Equal(p) {
			return i
		}
	}

	return -1
}
