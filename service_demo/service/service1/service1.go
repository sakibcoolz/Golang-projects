package service1

import (
	"gorm.io/gorm"
)

type Service1 struct {
}

type IService1 interface {
	Ping()
}

func GiveService1(DB *gorm.DB) (s1 Service1) {
	return s1
}

func (s Service1) Ping() {
}
