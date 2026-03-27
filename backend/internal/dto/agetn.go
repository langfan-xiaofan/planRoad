package dto

type ChatRequest struct {
	ConversationId uint     `json:"conversation_id"`
	UserID         uint     `json:"user_id"`
	Message        string   `json:"message"`
	Files          []string `json:"files"`
}

type ParseRequest struct {
	UserID uint   `json:"user_id"`
	File   string `json:"file"`
}
