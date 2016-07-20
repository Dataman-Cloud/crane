package account

import (
	"net/url"
)

type Group struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
}

type GroupFilter url.Values
