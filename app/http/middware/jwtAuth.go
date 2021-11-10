package middware

import (
	"ginapi/tools"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			tools.JsonReturnWithError(c, "auth in the request header is empty", 4001)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			tools.JsonReturnWithError(c, "the auth format in the request header is incorrect", 4001)
			c.Abort()
			return
		}
		mc, err := tools.ParseJwtToken(parts[1])
		if err != nil {
			tools.JsonReturnWithError(c, "invalid token", 4001)
			c.Abort()
			return
		}
		c.Set("jwt_uid", mc.UID)
		c.Next()
	}
}
