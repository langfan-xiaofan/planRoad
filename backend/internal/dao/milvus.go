package dao

var ColumnNames = []string{"id", "work_city", "position_type", "min_education", "core_skills", "priority", "salary_range", "company_name", "accept_intern", "need_work_exp", "industry", "job_embedding", "position_name"}

var Education = map[string]int{
	"不限":    0,
	"高中及以下": 1,
	"大专":    2,
	"本科":    3,
	"硕士":    4,
	"博士":    5,
}

var VectorWeight = map[string]float32{
	"skill_vector":          0.65,
	"position_vector":       0.15,
	"responsibility_vector": 0.15,
	"culture_vector":        0.05,
}

type JobInfo struct {
	JobID        int64    `json:"id"`
	WorkCity     string   `json:"work_city"`
	PositionType string   `json:"position_type"`
	MinEducation string   `json:"min_education"`
	CoreSkills   []string `json:"core_skills"`
	Priority     []string `json:"priority"`
	SalaryRange  string   `json:"salary_range"`
	CompanyName  string   `json:"company_name"`
	AcceptIntern bool     `json:"accept_intern"`
	NeedWorkExp  bool     `json:"need_work_exp"`
	Industry     string   `json:"industry"`
	JobVector
	PositionName string `json:"position_name"`
}

type JobVector struct {
	SkillVector          []float32 `json:"skill_vector"`
	PositionVector       []float32 `json:"position_vector"`
	ResponsibilityVector []float32 `json:"responsibility_vector"`
	CultureVector        []float32 `json:"culture_vector"`
}
