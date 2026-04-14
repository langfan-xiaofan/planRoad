package db

import (
	"backend/internal/config"

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
}
