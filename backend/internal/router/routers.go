package router

import (
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine, client *react.Agent) {
	initUserRouter(g)
	InitAgentRouter(g, client)
	initJobRouter(g)
}
