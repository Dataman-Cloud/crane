package account

type (
	ACL struct {
		Permission  string        `json:"Permission,omitempty"`
		Description string        `json:"Description,omitempty"`
		Rules       []*AccessRule `json:"Rules,omitempty"`
	}

	AccessRule struct {
		Path    string   `json:"Path,omitempty"`
		Methods []string `json:"Methods,omitempty"`
	}
)

func DefaultACLs() []*ACL {
	acls := []*ACL{}
	adminACL := &ACL{
		Permisson:   "admin",
		Description: "Administrator",
		Rules: []*AccessRule{
			{
				Path:    "*",
				Methods: []string{"*"},
			},
		},
	}
	acls = append(acls, adminACL)

	containersACLRO := &ACL{
		Permisson:   "containers:ro",
		Description: "Containers Read Only",
		Rules: []*AccessRule{
			{
				Path:    "/containers",
				Methods: []string{"GET"},
			},
		},
	}
	acls = append(acls, containersACLRO)

	containersACLRW := &ACL{
		Permisson:   "containers:rw",
		Description: "Containers",
		Rules: []*AccessRule{
			{
				Path:    "/containers",
				Methods: []string{"GET", "POST", "DELETE"},
			},
		},
	}
	acls = append(acls, containersACLRW)

	eventsACLRO := &ACL{
		Permisson:   "events:ro",
		Description: "Events Read Only",
		Rules: []*AccessRule{
			{
				Path:    "/api/events",
				Methods: []string{"GET"},
			},
		},
	}
	acls = append(acls, eventsACLRO)

	eventsACLRW := &ACL{
		Permisson:   "events:rw",
		Description: "Events",
		Rules: []*AccessRule{
			{
				Path:    "/api/events",
				Methods: []string{"GET", "POST", "DELETE"},
			},
		},
	}
	acls = append(acls, eventsACLRW)

	imagesACLRO := &ACL{
		Permisson:   "images:ro",
		Description: "Images Read Only",
		Rules: []*AccessRule{
			{
				Path:    "/images",
				Methods: []string{"GET"},
			},
		},
	}
	acls = append(acls, imagesACLRO)

	imagesACLRW := &ACL{
		Permisson:   "images:rw",
		Description: "Images",
		Rules: []*AccessRule{
			{
				Path:    "/images",
				Methods: []string{"GET", "POST", "DELETE"},
			},
		},
	}
	acls = append(acls, imagesACLRW)

	nodesACLRO := &ACL{
		Permisson:   "nodes:ro",
		Description: "Nodes Read Only",
		Rules: []*AccessRule{
			{
				Path:    "/api/nodes",
				Methods: []string{"GET"},
			},
		},
	}
	acls = append(acls, nodesACLRO)

	nodesACLRW := &ACL{
		Permisson:   "nodes:rw",
		Description: "Nodes",
		Rules: []*AccessRule{
			{
				Path:    "/api/nodes",
				Methods: []string{"GET", "POST", "DELETE"},
			},
		},
	}
	acls = append(acls, nodesACLRW)

	registriesACLRO := &ACL{
		Permisson:   "registries:ro",
		Description: "Registries Read Only",
		Rules: []*AccessRule{
			{
				Path:    "/api/registry",
				Methods: []string{"GET"},
			},
		},
	}
	acls = append(acls, registriesACLRO)

	registriesACLRW := &ACL{
		Permisson:   "registries:rw",
		Description: "Registries",
		Rules: []*AccessRule{
			{
				Path:    "/api/registry",
				Methods: []string{"GET", "POST", "DELETE"},
			},
		},
	}
	acls = append(acls, registriesACLRW)

	return acls
}

func DefaultPermissions() []string {
	acls := DefaultACLs{}
}
