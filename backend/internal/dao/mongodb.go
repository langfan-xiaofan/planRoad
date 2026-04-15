package dao

import (
	"backend/internal/db"
	"backend/internal/dto"
	"backend/internal/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetJob() ([]dto.GetJobRes, error) {
	var res []dto.GetJobRes
	collection := db.MongoDatabase.Collection("job_count")
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return res, errors.New("获取失败")
	}
	cursor.All(context.Background(), &res)
	if err := cursor.Err(); err != nil {
		return res, errors.New("获取失败")
	}
	return res, nil
}
func GetJobPicture(position string) (model.Position, error) { //获取职位的画像
	var res model.Position
	collection := db.MongoDatabase.Collection("picture2")
	filter := bson.M{"position": position}
	err := collection.FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return res, errors.New("未找到该职位")
	}
	return res, nil

}
func GetPositionSkills(position string) (dto.GetSkillsRes, error) { //获取职位的技能（用来打标签）
	var res dto.GetSkillsRes
	collection := db.MongoDatabase.Collection("picture2")
	filter := bson.M{"position": position}
	err := collection.FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return res, errors.New("未找到该职位")
	}
	return res, nil
}

func GetUpPath(position string) (dto.GetUpPathRes, error) { //获取晋升路径
	var res dto.GetUpPathRes
	collection := db.MongoDatabase.Collection("up_path")
	filter := bson.M{"position": position}
	err := collection.FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return res, errors.New("未找到该职位")
	}
	return res, nil
}
func GetAllJobPicture() ([]model.Position, error) {

	var res []model.Position

	collection := db.MongoDatabase.Collection("picture2")

	filter := bson.M{}

	cursor, err := collection.Find(context.Background(), filter)

	if err != nil {

		return res, err

	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {

		var pos model.Position

		if err := cursor.Decode(&pos); err != nil {

			// 如果某一行解码失败，可以选择跳过或直接返回错误

			return res, err

		}

		res = append(res, pos)

	}

	return res, nil

}
func GetJobAndPositon() ([]dto.GetPositionByJobRes, error) {
	var res []dto.GetPositionByJobRes
	collection := db.MongoDatabase.Collection("position")
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return res, errors.New("获取失败")
	}
	cursor.All(context.Background(), &res)
	if err := cursor.Err(); err != nil {
		return res, errors.New("获取失败")
	}
	return res, nil
}

func GetJobByPosition(position string) (dto.GetJobByPositionRes, error) { //通过职位获取大类
	var res dto.GetJobByPositionRes
	collection := db.MongoDatabase.Collection("picture2")
	filter := bson.M{"position": position}
	err := collection.FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return res, errors.New("未找到该职位")
	}
	return res, nil
}

func GetJobList() ([]dto.GetJobListRes, error) { //获取所有的Job
	var res []dto.GetJobListRes
	collection := db.MongoDatabase.Collection("job_count")
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return res, errors.New("获取失败")
	}
	cursor.All(context.Background(), &res)
	if err := cursor.Err(); err != nil {
		return res, errors.New("获取失败")
	}
	return res, nil
}
func CreatUserPicture(req dto.UserPictureReq) error {
	collection := db.MongoDatabase.Collection("picture1")
	filter := bson.M{"user_id": req.UserId}

	// 1. 先尝试更新
	update := bson.M{"$set": req}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	// 2. 如果 MatchedCount 为 0，说明数据库里没这条数据，那就执行插入
	if result.MatchedCount == 0 {
		_, err := collection.InsertOne(context.Background(), req)
		if err != nil {
			return err
		}
	}

	return nil
}
func GetUserPictureByUserId(userid uint) (dto.UserPictureRes, error) {
	var res dto.UserPictureRes
	collection := db.MongoDatabase.Collection("picture1")
	filter := bson.M{"user_id": userid}
	err := collection.FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return res, errors.New("未找到该职位")
	}
	return res, nil

}

func GetJobRelationship(req dto.GetPositionReq) (model.JobRelationship, error) {
	var res model.JobRelationship
	var err error
	collection := db.MongoDatabase.Collection("relationship")
	filter := bson.M{"origin_position": req.Position}
	err = collection.FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

//func GetPositionList() ([]dto.GetPositionReq, error) {
//	var res []model.Position
//	collection := db.MongoDatabase.Collection("position")
//	filter := bson.M{}
//	cursor, err := collection.Find(context.Background(), filter)
//	if err != nil {
//		return res, err
//	}
//	defer cursor.Close(context.Background())
//	for cursor.Next(context.Background()) {
//		var pos model.Position
//		if err := cursor.Decode(&pos); err != nil {
//			return res, err
//		}
//	}
//	if err := cursor.Err(); err != nil {
//		return res, err
//	}
//	return res, nil
//}
