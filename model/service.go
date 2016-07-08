package model

type Service struct {
	ID      uint64
	Name    string `gorm:"not null"`
	StackId uint64 `gorm:"type:varchar(64)"`
}
