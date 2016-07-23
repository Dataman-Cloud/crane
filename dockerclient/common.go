package dockerclient

const (
	PERMISSION_LABEL_PREFIX = "com.rolex.permissions"
)

type Permission struct {
	Perm    int    `json:"Perm"`
	Display string `json:"Display"`
}

type GroupPermission struct {
	Permission Permission `json:"Permission"`
	Group      string     `json:"Group"`
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
