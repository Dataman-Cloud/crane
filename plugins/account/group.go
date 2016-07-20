package account

import (
	"net/url"
)

type Group struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

type GroupFilter url.Values
