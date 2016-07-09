package model

import (
	"time"
)

type Stack struct {
	ID       uint64
	UserId   uint64    `gorm:"not null;type:varchar(64)"`
	Name     string    `gorm:"not null;type:varchar(64)"`
	Compose  string    `gorm:"type:text"`
	CreateAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
