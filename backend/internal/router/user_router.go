package router

import "github.com/gin-gonic/gin"

func initUserRouter(g *gin.Engine) {
	g.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "user router",
		})
	})
}
