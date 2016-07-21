package dockerclient

import (
	"strings"
)

type Perm string

var (
	PermReadOnly  Perm = "ro"
	PermReadWrite Perm = "rw"
	PermReadAdmin Perm = "admin"

	Perms []Perm = []Perm{PermReadOnly, PermReadWrite, PermReadAdmin}
)

// permissions on label
type Permission struct {
	Perm  Perm   `json:"Perm"`
	Group string `json:"Group"`
}

func PermValid(p Perm) bool {
	for _, perm := range Perms {
		if p == perm {
			return true
		}
	}

	return false
}

func (p Permission) Equal(o Permission) bool {
	return p.Perm == o.Perm && o.Group == p.Group
}

func PermissionsFromLabel(label string) []Permission {
	permissions := make([]Permission, 0)
	for _, str := range strings.SplitN(label, ",", -1) {
		pair := strings.SplitN(str, "-", 2)
		if len(pair) != 2 {
			continue
		}

		if !PermValid(Perm(pair[1])) {
			continue
		}
		permissions = append(permissions, Permission{Perm: Perm(pair[1]), Group: pair[0]})
	}

	return permissions
}

func PermissionsToLabel(permissions []Permission) string {
	pairs := make([]string, 0)
	for _, v := range permissions {
		pairs = append(pairs, strings.Join([]string{v.Group, string(v.Perm)}, "-"))
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
