package router

import (
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initUserRouter(g *gin.Engine) {
	userGroup := g.Group("/api/user")
	{
		userGroup.POST("/register", handler.RegisterHandler)
		userGroup.POST("/login", handler.LoginHandler)
	}
	authGroup := userGroup.Group("/")
	authGroup.Use(middleware.JwtMiddleware())
	{
		authGroup.POST("/avatar", handler.UploadAvatarHandler)
		authGroup.GET("/getuserpicture", handler.GetUserPictureHandler)
	}
}
