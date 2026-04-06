package router

import (
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

func InitAgentRouter(r *gin.Engine, client *react.Agent, fileParser *openai.Client, db *gorm.DB, redis *redis.Client) {
	agentHandler := handler.NewAgentHandler(client, fileParser, db, redis)
	agent := r.Group("/agent")
	agent.Use(middleware.JwtMiddleware())
	{
		agent.POST("/chat", agentHandler.Chat)
		agent.POST("/parse", agentHandler.Parse)
	}
}
