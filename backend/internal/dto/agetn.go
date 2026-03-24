package dto

type ChatRequest struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}
