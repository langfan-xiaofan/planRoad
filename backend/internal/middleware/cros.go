package middleware

import "github.com/gin-gonic/gin"

func CrosMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
		c.Writer.Header().Add("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type:text/event-stream")
		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		}
		c.Next()
	}
}
