package auth

import (
	"net/url"
)

type Group struct {
	ID   uint64 `json:"Id"`
	Name string `json:"Name"`
}

type GroupFilter url.Values
