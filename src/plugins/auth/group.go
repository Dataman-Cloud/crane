package auth

import (
	"net/url"
)

type Group struct {
	ID        uint64 `json:"Id"`
	Name      string `json:"Name" gorm:"not null"`
	CreaterId uint64 `json:"CreaterId"`
}

type GroupFilter url.Values
