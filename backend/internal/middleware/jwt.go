package middleware

import (
	"backend/pkg/res"
	"backend/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			res.Fail(c, 401, nil, "未携带token")
			c.Abort()
			return
		}
		authHeader = strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ParseToken(authHeader)
		if err != nil {
			res.Fail(c, 401, nil, "Token无效或已过期")
			c.Abort()
			return
		}
		c.Set("id", claims.Id)
		c.Next()
	}
}
