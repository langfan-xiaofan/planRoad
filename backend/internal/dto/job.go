package dto

type GetSkillsRes struct {
	Position string   `json:"position"`
	Skills   []string `json:"skills"`
}

type GetPositionReq struct {
	Position string `json:"position"`
}
type GetJobRes struct {
	Job   string `json:"job"`
	Count int    `json:"count"`
	Sum   int    `json:"sum"`
}
type GetUpPathRes struct {
	Position     string   `bson:"position"`
	CareerPath   []string `bson:"career_path"`
	GrowthAdvice string   `bson:"growth_advice"`
	Description  string   `bson:"job_description"`
}

type GetPositionByJobRes struct {
	Job      string   `json:"job"`
	Position []string `json:"position"`
}

type GetJobByPositionRes struct {
	Job      string `json:"job"`
	Position string `json:"position"`
}

type GetJobListRes struct {
	Job string `json:"job"`
}
