package middware

import (
	"ginapi/app/services"
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
			response.Message(c, "auth in the request header is empty", services.AuthErrCode)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Message(c, "the auth format in the request header is incorrect", services.AuthErrCode)
			c.Abort()
			return
		}
		mc, err := tools.ParseJwtToken(parts[1])
		if err != nil {
			response.Message(c, "invalid token", services.AuthErrCode)
			c.Abort()
			return
		}
		c.Set("jwt_uid", mc.UID)
		c.Next()
	}
}
