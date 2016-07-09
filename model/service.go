package model

type Service struct {
	ID      uint64
	UserId  uint64 `gorm:"not null;type:varchar(64)"`
	Name    string `gorm:"not null"`
	StackId uint64 `gorm:"type:varchar(64)"`
}
