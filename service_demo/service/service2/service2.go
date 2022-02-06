package service2

import (
	"service_demo/dao/dao2"

	"gorm.io/gorm"
)

type Service2 struct {
	Curd dao2.IDao1
}

type IService2 interface {
	Ping()
}

func GiveService2(DB *gorm.DB) (s2 Service2) {
	s2 = Service2{
		Curd: dao2.GiveDao2(DB),
	}
	return s2
}

func (s Service2) Ping() {
	s.Curd.DatabasePing()
}
