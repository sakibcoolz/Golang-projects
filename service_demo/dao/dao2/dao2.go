package dao2

import (
	"log"

	"gorm.io/gorm"
)

type IDao1 interface {
	DatabasePing()
}

type DataStore struct {
	DB *gorm.DB
}

func GiveDao2(DB *gorm.DB) (DS DataStore) {
	DS = DataStore{DB: DB}
	return DS
}

func (DS DataStore) DatabasePing() {
	err := DS.DB.Raw("").Error
	if err != nil {
		log.Println(err)
	}
}
