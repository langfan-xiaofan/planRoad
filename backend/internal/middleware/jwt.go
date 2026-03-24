package middleware

import (
	"backend/pkg/res"
	"backend/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			res.Fail(c, 401, nil, "未携带token")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			res.Fail(c, 401, nil, "token错误")
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			res.Fail(c, 401, nil, "token无效")
			c.Abort()
			return
		}
		c.Set("userid", claims.Id)
		c.Set("resumeid", claims.RemuseId)
		c.Next()
	}
}
