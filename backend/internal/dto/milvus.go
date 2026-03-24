package dto

type BasicInfo struct {
	PositionName string `json:"position_name"` // 岗位名称
	CompanyName  string `json:"company_name"`  // 公司名称
	PositionType string `json:"position_type"` // 岗位类型（技术/销售/运营等）
	WorkCity     string `json:"work_city"`     // 工作城市
	WorkAddress  string `json:"work_address"`  // 详细工作地址
	RecruitNum   string `json:"recruit_num"`   // 招聘人数
	PublishTime  string `json:"publish_time"`  // 发布时间（格式：YYYY-MM-DD）
}

// 学历要求
type Education struct {
	MinEducation    string   `json:"min_education"`     // 最低学历
	MajorLimit      []string `json:"major_limit"`       // 限定专业
	AcceptNonMajor  bool     `json:"accept_non_major"`  // 是否接受非相关专业
	AcceptFreshGrad int      `json:"accept_fresh_grad"` // 是否接受应届生（1=是，0=否）
	AcceptIntern    int      `json:"accept_intern"`     // 是否接受实习生（1=是，0=否）
}

// 经验要求
type Experience struct {
	MinWorkYears        string   `json:"min_work_years"`        // 最低工作年限
	NeedWorkExp         bool     `json:"need_work_exp"`         // 是否需要工作经验
	IndustryExpRequired []string `json:"industry_exp_required"` // 要求的行业经验
}

// 技能要求
type SkillRequire struct {
	CoreSkills         []string `json:"core_skills"`         // 核心技能
	SoftSkills         []string `json:"soft_skills"`         // 软技能
	CertificateRequire []string `json:"certificate_require"` // 必备证书
	LanguageRequire    []string `json:"language_require"`    // 语言要求
}

// 任职要求（汇总）
type Requirement struct {
	Education    Education    `json:"education"`     // 学历要求
	Experience   Experience   `json:"experience"`    // 经验要求
	SkillRequire SkillRequire `json:"skill_require"` // 技能要求
	OtherRequire []string     `json:"other_require"` // 其他要求
}

// 工作职责
type Responsibility struct {
	CoreWork          []string `json:"core_work"`          // 核心工作
	DailyWork         []string `json:"daily_work"`         // 日常工作
	PerformanceTarget []string `json:"performance_target"` // 业绩目标
}

// 薪资福利
type WelfareInfo struct {
	SalaryRange  string   `json:"salary_range"`  // 薪资范围
	SalaryType   string   `json:"salary_type"`   // 薪资类型
	Welfare      []string `json:"welfare"`       // 基础福利
	WorkingHours string   `json:"working_hours"` // 工作时间
	OtherWelfare []string `json:"other_welfare"` // 其他福利
}

// 公司信息
type CompanyInfo struct {
	CompanyScale  string `json:"company_scale"`  // 公司规模（小/中/大）
	Industry      string `json:"industry"`       // 所属行业
	CompanyNature string `json:"company_nature"` // 公司性质
	CompanyIntro  string `json:"company_intro"`  // 公司简介
}

// 优先条件
type Priority struct {
	PreferSkills     []string `json:"prefer_skills"`     // 优先技能
	PreferExperience []string `json:"prefer_experience"` // 优先经验
	PreferCharacter  []string `json:"prefer_character"`  // 优先特质
}

// 顶层结构体（对应完整JSON）
type JobInfo struct {
	BasicInfo      BasicInfo      `json:"basic_info"`
	Requirement    Requirement    `json:"requirement"`
	Responsibility Responsibility `json:"responsibility"`
	WelfareInfo    WelfareInfo    `json:"welfare_info"`
	CompanyInfo    CompanyInfo    `json:"company_info"`
	Priority       Priority       `json:"priority"`
}
