package account

type Role struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

func (r *Role) Parent() *Role {
	return nil
}

func (r *Role) Children() []*Role {
	roles := make([]*Role, 0)
	return roles
}

func (r *Role) Permissions() []string {
	return []string{}
}
