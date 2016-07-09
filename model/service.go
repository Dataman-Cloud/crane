package model

import (
	"time"
)

type Service struct {
	ID         uint64
	UserId     uint64    `gorm:"not null;type:varchar(64)"`
	Name       string    `gorm:"not null"`
	StackId    uint64    `gorm:"type:varchar(64)"`
	CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
