package controller2

import (
	"service_demo/service/service1"

	"gorm.io/gorm"
)

type Controller2 struct {
	service1 service1.IService1
}

type IController2 interface {
	Ping()
}

func GiveController2(DB *gorm.DB) (c2 Controller2) {
	c2 = Controller2{
		service1: service1.GiveService1(DB),
	}
	return c2
}

func (c Controller2) Ping() {
}
