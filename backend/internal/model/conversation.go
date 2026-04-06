package model

import (
	"time"

	"github.com/cloudwego/eino/schema"
)

// Conversation 对话
type Conversation struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Message 消息
type Message struct {
	Id             uint            `json:"id" gorm:"primary_key"`
	ConversationId uint            `json:"conversation_id"`
	UserId         uint            `json:"user_id"`
	Role           schema.RoleType `json:"role"`
	Content        string          `json:"content"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}
