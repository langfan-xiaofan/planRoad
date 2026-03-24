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
		// &model.Article{}, // 如果以后有其他表，继续往后加
	)
	if err != nil {
		return err // 迁移失败通常是因为权限不足或数据库连接中断
	}
	return nil
}
