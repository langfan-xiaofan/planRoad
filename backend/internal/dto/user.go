package dto

import (
	"mime/multipart"
)

type RegisterReq struct {
	Username string `json:"username"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type AvatarReq struct {
	Avatar *multipart.FileHeader `form:"avatar"`
}
type AvatarRes struct {
	AvatarUrl string `json:"avatarurl"`
}
type UserPictureReq struct {
	UserId          uint     `bson:"user_id" json:"user_id" `
	Name            string   `bson:"name" json:"name"`                         //名字
	School          string   `bson:"school" json:"school"`                     //学历
	Job             string   `bson:"job" json:"job"`                           //推荐的工作
	Position        string   `bson:"position" json:"position"`                 //推荐的职位
	Skills          []string `bson:"skills" json:"skills"`                     //会的技能
	Certifications  []string `bson:"certifications" json:"certifications"`     //证书
	Innovation      string   `bson:"innovation" json:"innovation"`             //创新能力
	LearningAbility string   `bson:"learning_ability" json:"learning_ability"` //学习能力
	StressTolerance string   `bson:"stress_tolerance" json:"stressTolerance"`  //抗压能力
	Communication   string   `bson:"communication" json:"communication"`       //沟通能力
	Experience      string   `bson:"experience" json:"experience"`             //项目经验
	PublicBase      []string `bson:"publicbase" json:"publicBase"`             //基础
}

type UserPictureRes struct {
	UserId uint `bson:"user_id" json:"user_id"`
	//ID              primitive.ObjectID `json:"id" bson:"_id"`    //id(会存到user里面)
	Name            string   `bson:"name" json:"name"`                         //名字
	School          string   `bson:"school" json:"school"`                     //学历
	Job             string   `bson:"job" json:"job"`                           //推荐的工作
	Position        string   `bson:"position" json:"position"`                 //推荐的职位
	Skills          []string `bson:"skills" json:"skills"`                     //会的技能
	Certifications  []string `bson:"certifications" json:"certifications"`     //证书
	Innovation      string   `bson:"innovation" json:"innovation"`             //创新能力
	LearningAbility string   `bson:"learning_ability" json:"learning_ability"` //学习能力
	StressTolerance string   `bson:"stress_tolerance" json:"stress_tolerance"` //抗压能力
	Communication   string   `bson:"communication" json:"communication"`       //沟通能力
	Experience      string   `bson:"experience" json:"experience"`             //项目经验
	PublicBase      []string `bson:"publicbase" json:"publicbase"`             //基础
}
