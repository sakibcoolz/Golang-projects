package service

import (
	"service_demo/service/service1"
	"service_demo/service/service2"

	"gorm.io/gorm"
)

type Service struct {
	service1.IService1
	service2.IService2
}

func GiveService(DB *gorm.DB) (s Service) {
	return Service{
		IService1: service1.GiveService1(DB),
		IService2: service2.GiveService2(DB),
	}
}
