package middware

import (
	"ginapi/tools"
	"ginapi/tools/response"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Msg(c, "invalid token", 401)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Msg(c, "invalid token", 401)
			c.Abort()
			return
		}
		mc, err := tools.ParseJwtToken(parts[1])
		if err != nil {
			response.Msg(c, "invalid token", 401)
			c.Abort()
			return
		}
		c.Set("jwt_uid", mc.UID)
		c.Next()
	}
}
