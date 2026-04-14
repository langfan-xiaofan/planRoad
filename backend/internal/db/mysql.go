package db

import (
	"backend/internal/config"
	"backend/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDatabase *gorm.DB

func MysqlInit() error {
	var err error
	MysqlDatabase, err = gorm.Open(mysql.Open(config.Conf.Dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = MysqlDatabase.AutoMigrate(
		&model.User{},
		&model.Message{},
	)
	if err != nil {
		return err
	}
	return nil
}
