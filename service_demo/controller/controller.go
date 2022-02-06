package controller

import (
	"service_demo/controller/controller1"
	"service_demo/controller/controller2"

	"gorm.io/gorm"
)

type Controller struct {
	IControl1 controller1.IController1
	IControl2 controller2.IController2
}

func GiveController(DB *gorm.DB) (c Controller) {
	c = Controller{
		IControl1: controller1.GiveController1(DB),
		IControl2: controller2.GiveController2(DB),
	}
	return
}
