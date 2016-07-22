package dockerclient

const (
	PERMISSION_LABEL_PREFIX = "com.permissions"
)

type Permission struct {
	Perm    int    `json:"Perm"`
	Display string `json:"Display"`
	Group   string `json:"Group"`
}

var (
	PermReadOnly  Permission = Permission{Perm: 0, Display: "r", Group: "foo"}
	PermReadWrite Permission = Permission{Perm: 1, Display: "w", Group: "foo"}
	PermReadAdmin Permission = Permission{Perm: 2, Display: "x", Group: "foo"}

	Perms []Permission = []Permission{PermReadOnly, PermReadWrite, PermReadAdmin}
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
