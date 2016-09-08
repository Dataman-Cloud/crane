package registry

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Notification holds all events.
type Notification struct {
	Events []Event
}

// Event holds the details of a event.
type Event struct {
	ID        string `json:"Id"`
	TimeStamp time.Time
	Action    string
	Target    *Target
	Request   *Request
	Actor     *Actor
}

// Target holds information about the target of a event.
type Target struct {
	MediaType  string
	Digest     string
	Repository string
	URL        string `json:"Url"`
}

// Actor holds information about actor.
type Actor struct {
	Name string
}

// Request holds information about a request.
type Request struct {
	ID        string `json:"Id"`
	Method    string
	UserAgent string
}

type Image struct {
	gorm.Model

	LatestTag string `json:"LatestTag"`
	Namespace string `json:"Namespace" gorm:"not null"`
	Image     string `json:"Image" gorm:"not null"`
	Publicity uint8  `json:"Publicity" gorm:"not null default 0"`

	PushCount int64 `json:"PushCount" gorm:"-"`
	PullCount int64 `json:"PullCount" gorm:"-"`
}

type Tag struct {
	gorm.Model

	Digest    string `json:"Digest"`
	Tag       string `json:"Tag"`
	Namespace string `json:"Namespace" gorm:"not null"`
	Image     string `json:"Image" gorm:"not null"`
	Size      uint64 `json:"Size"`

	PushCount int64 `json:"PushCount" gorm:"-"`
	PullCount int64 `json:"PullCount" gorm:"-"`
}

type ImageAccess struct {
	gorm.Model

	Namespace    string `json:"Namespace" gorm:"not null"`
	Image        string `json:"Image" gorm:"not null"`
	Digest       string `json:"Digest"`
	AccountEmail string `json:"AccountEmail"`
	Action       string `json:"Action"`
}

type V2RegistryResponse struct {
	SchemaVersion int                      `json:"schemaVersion"`
	MediaType     string                   `json:"mediaType"`
	Config        V2RegistryResponseConfig `json:"config"`
	Layers        V2RegistryResponseLayers `json:"layers"`
}

type V2RegistryResponseConfig struct {
	MediaType string `json:"mediaType"`
	Size      uint64 `json:"size"`
	Digest    string `json:"digest"`
}

type V2RegistryResponseLayers []V2RegistryResponseLayer

type V2RegistryResponseLayer struct {
	MediaType string `json:"mediaType"`
	Size      uint64 `json:"size"`
	Digest    string `json:"digest"`
}
