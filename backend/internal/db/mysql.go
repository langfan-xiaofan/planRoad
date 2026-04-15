package db

import (
	"backend/internal/config"
	"backend/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDatabase *gorm.DB

func MysqlInit() {
	var err error
	MysqlDatabase, err = gorm.Open(mysql.Open(config.Conf.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = MysqlDatabase.AutoMigrate(&model.Message{})
	if err != nil {
		panic(err)
	}

	//MysqlDatabase.AutoMigrate(&model.{})
}
