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
