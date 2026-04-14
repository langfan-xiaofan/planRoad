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

type PositionAndAdvice struct {
	PositionName string `bson:"position_name" json:"position_name"`
	Advice       string `bson:"advice" json:"advice"`
}
type JobRelationship struct {
	OriginPosition string              `bson:"origin_position" json:"origin_position"`
	Suggestion     []PositionAndAdvice `bson:"suggestion" json:"suggestion"`
}
