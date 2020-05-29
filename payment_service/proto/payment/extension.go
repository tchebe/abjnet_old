package payment

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//BeforeCreate is a postgresql gorm related function enabling us to
//add uuid for example
func (model *Payment) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}
