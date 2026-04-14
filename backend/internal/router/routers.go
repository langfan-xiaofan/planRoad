package router

import (
	eino "github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

func InitRouter(g *gin.Engine, client *react.Agent, fileParser *openai.Client, db *gorm.DB, redis *redis.Client, chatModel *eino.ChatModel) {
	initUserRouter(g)
	InitAgentRouter(g, client, fileParser, db, redis, chatModel)
	initJobRouter(g)
}
