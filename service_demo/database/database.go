package database

import (
	"fmt"
	"service_demo/utils/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataStore struct {
	DB *gorm.DB
}

func DbConnector() (DB *gorm.DB, err error) {
	config := config.ReadConfig()
	dsn := fmt.Sprint(config.Database.User + ":" + config.Database.Password + "@" + config.Database.ConnectionType + "(" + config.Database.Host + ":" + config.Database.Port + ")/" + config.Database.Schema + "?" + config.Database.Charset)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		sqlDB, _ := DB.DB()
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	return DB, err
}
