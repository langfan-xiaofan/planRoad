package router

import (
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

func InitRouter(g *gin.Engine, client *react.Agent, fileParser *openai.Client, db *gorm.DB, redis *redis.Client) {
	initUserRouter(g)
	InitAgentRouter(g, client, fileParser, db, redis)
	initJobRouter(g)
}
