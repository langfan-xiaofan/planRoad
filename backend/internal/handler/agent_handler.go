package handler

import (
	"backend/internal/dto"
	"backend/internal/service"
	"fmt"

	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/gin-gonic/gin"
)

type AgentHandler struct {
	AgentService *service.AgentService
}

func NewAgentHandler(client *react.Agent) *AgentHandler {
	return &AgentHandler{
		AgentService: service.NewAgentService(client),
	}
}

func (h *AgentHandler) Chat(c *gin.Context) {
	var req dto.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var message []*schema.Message
	message = append(message, &schema.Message{
		Role:    schema.User,
		Content: req.Message,
	})
	fmt.Println(req)
	if err := h.AgentService.Chat(&req, message, c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}
