package controller1

import (
	"net/http"
	"service_demo/service/service1"
	"service_demo/service/service2"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller1 struct {
	service1 service1.IService1
	service2 service2.IService2
}

type IController1 interface {
	Ping(c *gin.Context)
}

func GiveController1(DB *gorm.DB) (c1 Controller1) {
	c1 = Controller1{
		service1: service1.GiveService1(DB),
		service2: service2.GiveService2(DB),
	}
	return c1
}

func (co Controller1) Ping(c *gin.Context) {
	co.service1.Ping()
	c.String(http.StatusAccepted, "sakib mulla")
}
