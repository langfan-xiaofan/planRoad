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

//	func CreatUserPicture(req dto.UserPictureReq, id uint) error {
//		collection := db.MongoDatabase.Collection("picture1")
//
//		// 插入数据
//		insertResult, err := collection.InsertOne(context.Background(), req)
//		if err != nil {
//			return errors.New("MongoDB 插入失败")
//		}
//
//		var mongoIDStr string
//		// 使用 type switch 兼容多种可能的返回类型
//		switch v := insertResult.InsertedID.(type) {
//		case primitive.ObjectID:
//			mongoIDStr = v.Hex()
//		case *primitive.ObjectID:
//			mongoIDStr = v.Hex()
//		default:
//			// 如果断言都失败了，说明驱动返回了奇怪的类型，用格式化拿到它
//			mongoIDStr = fmt.Sprintf("%v", v)
//			// 兜底处理：去掉可能存在的包装字符
//			mongoIDStr = strings.TrimPrefix(mongoIDStr, "ObjectID(\"")
//			mongoIDStr = strings.TrimSuffix(mongoIDStr, "\")")
//		}
//
//		// 严谨校验：如果拿到的还是全0或空，直接拦截
//		if mongoIDStr == "" || mongoIDStr == "000000000000000000000000" {
//			return errors.New("生成的 MongoDB ID 无效")
//		}
//
//		// 更新到 MySQL
//		err = UpdateResumeId(mongoIDStr, id)
//		if err != nil {
//			return fmt.Errorf("MySQL 更新失败: %v", err)
//		}
//
//		fmt.Println("成功创建画像并关联 MySQL，ID:", mongoIDStr)
//		return nil
//	}
func CreatUserPicture(req dto.UserPictureReq) error {
	collection := db.MongoDatabase.Collection("picture1")
	_, err := collection.InsertOne(context.Background(), req)
	if err != nil {
		return errors.New("MongoDB 插入失败")
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
