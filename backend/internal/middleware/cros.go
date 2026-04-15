package middleware

import (
	"github.com/gin-gonic/gin"
)

func CrosMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//origin := c.Request.Header.Get("Origin")
		//fmt.Println(origin)
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,ngrok-skip-browser-warning")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type:text/event-stream")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		}
		c.Next()
	}
}
