package dao

import (
	"service_demo/dao/dao2"

	"gorm.io/gorm"
)

type CurdOperation struct {
	dao2.IDao1
}

func GiveCurdOperation(DB *gorm.DB) (co CurdOperation) {
	co = CurdOperation{IDao1: dao2.GiveDao2(DB)}
	return co
}
