package dao

import (
	"backend/internal/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/cloudwego/eino/schema"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ConversationDao struct {
	db    *gorm.DB
	Redis *redis.Client
}

func NewConversationDao(db *gorm.DB, redis *redis.Client) *ConversationDao {
	return &ConversationDao{
		db:    db,
		Redis: redis,
	}
}

func (d *ConversationDao) AddMessageToMysql(message *model.Message) error {
	return d.db.Table("messages").Create(message).Error
}

func (d *ConversationDao) GetMessagesFromMysql(conversationId uint) ([]*schema.Message, error) {
	var messages []*model.Message
	err := d.db.Table("messages").Where("user_id = ?", conversationId).Order("created_at asc").Scan(&messages).Limit(20).Error
	if err != nil {
		return nil, err
	}
	// 转换为 schema.Message 类型
	schemaMessages := make([]*schema.Message, 0, len(messages))
	for _, msg := range messages {
		schemaMessages = append(schemaMessages, &schema.Message{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}
	fmt.Println("历史对话:", schemaMessages)
	return schemaMessages, nil
}

func (d *ConversationDao) GetAllHistoryMessageByConversationId(conversationId uint) ([]*model.Message, error) {
	var messages []*model.Message
	err := d.db.Where("conversation_id = ?", conversationId).Order("created_at asc").Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (d *ConversationDao) GetMessagesFromRedis(conversationId uint) ([]*schema.Message, error) {
	var messages []*model.Message
	msgStr, err := d.Redis.LRange(context.Background(), "conversation:"+strconv.Itoa(int(conversationId)), 0, -1).Result()
	if err != nil {
		return make([]*schema.Message, 0), err
	}
	for _, msg := range msgStr {
		var message model.Message
		err := json.Unmarshal([]byte(msg), &message)
		if err != nil {
			return make([]*schema.Message, 0), err
		}
		messages = append(messages, &message)
	}
	// 转换为 schema.Message 类型
	schemaMessages := make([]*schema.Message, 0, len(messages))
	for _, msg := range messages {
		schemaMessages = append(schemaMessages, &schema.Message{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}
	if len(schemaMessages) == 0 {
		return make([]*schema.Message, 0), errors.New("no messages found")
	}
	fmt.Println("Redis 对话:", schemaMessages)
	return schemaMessages, nil
}

func (d *ConversationDao) AddMessageToRedis(conversationId uint, message *model.Message) error {
	jsonStr, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = d.Redis.RPush(context.Background(), "conversation:"+strconv.Itoa(int(conversationId)), string(jsonStr)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (d *ConversationDao) SummaryMessage(conversationId uint, message *[]model.Message) error {
	err := d.Redis.Del(context.Background(), "conversation:"+strconv.Itoa(int(conversationId))).Err()
	if err != nil {
		return err
	}
	for _, msg := range *message {
		err = d.AddMessageToRedis(conversationId, &msg)
	}
	if err != nil {
		return err
	}
	return nil
}
