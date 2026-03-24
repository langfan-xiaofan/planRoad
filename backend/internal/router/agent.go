package router

import (
	"backend/internal/handler"

	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/gin-gonic/gin"
)

func InitAgentRouter(r *gin.Engine, client *react.Agent) {
	agentHandler := handler.NewAgentHandler(client)
	agent := r.Group("/agent")
	{
		agent.POST("/chat", agentHandler.Chat)
	}
}
