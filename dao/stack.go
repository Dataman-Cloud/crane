package dao

import (
	"github.com/Dataman-Cloud/rolex/model"
	"github.com/Dataman-Cloud/rolex/util/db"
)

func CreateStack(stack *model.Stack) error {
	return db.DB().Table("stack").Create(stack).Error
}

func UpdateStackById(stack *model.Stack) error {
	return db.DB().Table("stack").Save(stack).Error
}
