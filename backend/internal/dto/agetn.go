package dto

type ChatRequest struct {
	UserID  uint   `json:"user_id"`
	Message string `json:"message"`
}

type ParseRequest struct {
	UserID uint   `json:"user_id"`
	File   string `json:"file"`
}

type GetPositionDifferenceRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type InsightRequest struct {
}

type ResumeRadarReq struct {
	UserID uint   `json:"user_id"`
	Job    string `json:"job"`
}

type ResumeRadarRes struct {
	Skills         uint `json:"专业技能"`
	Certificate    uint `json:"证书要求"`
	Learning       uint `json:"创新能力"`
	Innovation     uint `json:"学习能力"`
	Tolerance      uint `json:"抗压能力"`
	Communications uint `json:"沟通能力"`
	Experience     uint `json:"实习能力"`
}
