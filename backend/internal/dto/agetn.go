package dto

type ChatRequest struct {
	UserID  uint   `json:"user_id"`
	Message string `json:"message"`
}

type ParseRequest struct {
	UserID uint   `json:"user_id"`
	File   string `json:"file"`
}
