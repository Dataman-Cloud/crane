package model

type Stack struct {
	ID      uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	UserId  uint64 `gorm:"not null;type:varchar(64)"`
	Name    string `gorm:"not null;type:varchar(64)"`
	Compose string `gorm:"type:text"`
}
