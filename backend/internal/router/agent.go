package router

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	eino "github.com/cloudwego/eino-ext/components/model/openai"

	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

func InitAgentRouter(r *gin.Engine, client *react.Agent, fileParser *openai.Client, db *gorm.DB, redis *redis.Client, chatModel *eino.ChatModel) {
	agentHandler := handler.NewAgentHandler(client, fileParser, db, redis, chatModel)
	agent := r.Group("/api")
	r.POST("agent/chatV2", agentHandler.ChatV2)
	agent.POST("/resume/parse", agentHandler.Parse)
	agent.Use(middleware.JwtMiddleware())
	{
		agent.POST("/chat", agentHandler.Chat)
	}
}
