package router

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func initUserRouter(g *gin.Engine) {
	userGroup := g.Group("/user")
	{
		userGroup.POST("/register", handler.RegisterHandler)
		userGroup.POST("/login", handler.LoginHandler)
	}
	//userGroup.Use(middleware.JwtMiddleware())

}
