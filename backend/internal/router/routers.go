package router

import "github.com/gin-gonic/gin"

func InitRouter(g *gin.Engine) {
	initUserRouter(g)
}
