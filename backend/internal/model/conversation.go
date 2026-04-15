package model

import (
	"time"

	"github.com/cloudwego/eino/schema"
)

// Conversation 对话

// Message 消息
type Message struct {
	Id        uint            `json:"id" gorm:"primaryKey"`
	UserId    uint            `json:"user_id"`
	Role      schema.RoleType `json:"role"`
	Content   string          `json:"content"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
