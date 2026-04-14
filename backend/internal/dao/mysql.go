package dao

import (
	"backend/internal/db"
)

func UpdateResumeId(resumeId string, id uint) error {
	return db.MysqlDatabase.Table("users").Where("id = ?", id).Update("resume_id", resumeId).Error
}

//func GetUserPicture(id uint) (dto.UserPictureRes, error) { //获取用户画像
//	var res dto.UserPictureRes
//	var resumeId string
//	err := db.MysqlDatabase.Table("users").Where("id = ?", id).Select("resume_id").Scan(&resumeId).Error
//	fmt.Println(resumeId)
//	res, err = GetUserPictureById(resumeId)
//	if err != nil {
//		return res, err
//	}
//	return res, nil
//
//}
