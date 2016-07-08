package model

type Stack struct {
	ID      uint64
	Name    string `gorm:"not null;type:varchar(64)"`
	Compose string `gorm:"type:text"`
}
