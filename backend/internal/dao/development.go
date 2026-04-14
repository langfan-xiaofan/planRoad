package dao

import "backend/internal/db"

func SetStartPosition(UserId uint, Position string) error {
	var count int64
	db.MysqlDatabase.Table("development_node").Where("user_id=?", UserId).Count(&count)
	if count > 0 {
		err := db.MysqlDatabase.Table("development_node").Where("user_id=?", UserId).Update("position", Position).Error
		if err != nil {
			return err
		}
	} else {
		err := db.MysqlDatabase.Table("development_node").Where("user_id=?", UserId).Create(Position).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func GetPlan(UserId uint, StartPosition string) {
	
}
