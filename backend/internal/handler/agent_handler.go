package handler

import (
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/service"
	"fmt"

	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

type AgentHandler struct {
	AgentService *service.AgentService
}

func NewAgentHandler(client *react.Agent, fileParser *openai.Client, db *gorm.DB, redis *redis.Client) *AgentHandler {
	return &AgentHandler{
		AgentService: service.NewAgentService(client, fileParser, db, redis),
	}
}
func (h *AgentHandler) Chat(c *gin.Context) {
	var req dto.ChatRequest
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req.UserID = uint(c.GetInt("user_id"))
	req.ConversationId = uint(c.GetInt("conversation_id"))
	file := form.File["file"]
	filePaths := make([]string, 0)
	for _, f := range file {
		if err := c.SaveUploadedFile(f, fmt.Sprintf("./uploads/%s", f.Filename)); err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		filePaths = append(filePaths, fmt.Sprintf("./uploads/%s", f.Filename))
	}
	if !h.AgentService.ConversationDao.ExistConversation(req.ConversationId) {
		h.AgentService.ConversationDao.AddConversationToMysql(&model.Conversation{
			UserId: uint(c.GetInt("userid")),
		})
		return
	}
	// 若没有消息，默认分析文件
	req.Message = form.Value["message"][0]
	if req.Message == "" {
		req.Message = "请分析这个文件"
	}
	// 上传文件并获取文件ID
	fileIDs, err := h.AgentService.UploadFile(filePaths, h.AgentService.FileParser)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// 解析文件返回文件内容
	parsedFiles, err := h.AgentService.ParseFile(fileIDs, h.AgentService.FileParser, req.Message, c)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// 添加解析后的文件内容到Redis和数据库
	h.AgentService.ConversationDao.AddMessageToRedis(req.ConversationId, &model.Message{
		ConversationId: req.ConversationId,
		UserId:         req.UserID,
		Role:           schema.Assistant,
		Content:        parsedFiles.Content,
	})
	h.AgentService.ConversationDao.AddMessageToMysql(&model.Message{
		ConversationId: req.ConversationId,
		UserId:         req.UserID,
		Role:           schema.Assistant,
		Content:        parsedFiles.Content,
	})
	// 从Redis获取历史消息，若没有则从数据库获取,若数据库也没有则返回空消息
	var messages []*schema.Message
	messages, err = h.AgentService.ConversationDao.GetMessagesFromRedis(req.ConversationId)
	if err != nil {
		messages, err = h.AgentService.ConversationDao.GetMessagesFromMysql(req.ConversationId)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
	defer func() {
		if len(messages) > 20 {
			messages = append(messages, &schema.Message{
				Role:    schema.User,
				Content: "请总结一下以上内容",
			})
			newMess, err := h.AgentService.Chat(&req, messages, c)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			h.AgentService.ConversationDao.SummaryMessage(req.ConversationId, &model.Message{
				ConversationId: req.ConversationId,
				UserId:         req.UserID,
				Role:           schema.Assistant,
				Content:        newMess,
			})
		}
	}()
	messages = append(messages, &schema.Message{
		Role:    schema.User,
		Content: req.Message,
	})
	h.AgentService.ConversationDao.AddMessageToMysql(&model.Message{
		ConversationId: req.ConversationId,
		UserId:         req.UserID,
		Role:           schema.User,
		Content:        req.Message,
	})
	if _, err = h.AgentService.Chat(&req, messages, c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (h *AgentHandler) Parse(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	file := form.File["file"]
	filePaths := make([]string, 0)
	for _, f := range file {
		if err := c.SaveUploadedFile(f, fmt.Sprintf("./uploads/%s", f.Filename)); err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		filePaths = append(filePaths, fmt.Sprintf("./uploads/%s", f.Filename))
	}
	userMessage := form.Value["message"][0]
	if userMessage == "" {
		userMessage = "请分析这个文件"
	}
	fileIDs, err := h.AgentService.UploadFile(filePaths, h.AgentService.FileParser)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	parsedFiles, err := h.AgentService.ParseFile(fileIDs, h.AgentService.FileParser, userMessage, c)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"parsedFiles": parsedFiles})
}
