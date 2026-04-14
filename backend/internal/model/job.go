package model

type Position struct {
	Id              string   `bson:"id"`
	Job             string   `bson:"job"`
	Position        string   `bson:"position"`
	Skills          []string `bson:"skills"`
	Certifications  []string `bson:"certifications"`
	Innovation      string   `bson:"innovation"`
	LearningAbility string   `bson:"learning_ability"`
	StressTolerance string   `bson:"stress_tolerance"`
	Communication   string   `bson:"communication"`
	ExperienceReq   string   `bson:"experience_req"`
	PublicBase      []string `bson:"publicbase"`
}

type UserPosition struct {
	UserId   string `json:"user_id"`
	Position string `json:"position"`
}

type DevelopmentNode struct {
	UserId       uint   `gorm:"user_id"`
	Position     string `gorm:"position"`
	NextPosition string `gorm:"next_position"`
	Plan         string `gorm:"plan"`
}

type Plan struct {
	UserId        uint   `gorm:"user_id"`
	Position      string `gorm:"position"`
	KnowledgeName string `gorm:"knowledge_name"`
	Time          string `gorm:"time"`
	Skills        string `gorm:"skills"`
}
