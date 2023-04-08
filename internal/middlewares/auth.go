package middlewares

import (
	"github.com/ayanami-desu/iropic-go/api/response"
	"github.com/ayanami-desu/iropic-go/internal/auth"
	"github.com/gin-gonic/gin"
	"strings"
)

func Auth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.Set("user", "guest")
		c.Next()
		return
	}
	if !strings.HasPrefix(token, "Bearer") {
		response.ErrorStrResp(c, "不支持的认证方式", 400)
		return
	}
	token = strings.Split(token, " ")[1]
	userClaims, err := auth.ParseToken(token)
	if err != nil {
		response.ErrorResp(c, err, 401)
		c.Abort()
		return
	}
	if userClaims.Username == "admin" {
		c.Set("user", "admin")
		c.Next()
		return
	}
	response.ErrorStrResp(c, "不支持的用户", 401)
	c.Abort()
	return
}

func AuthAdmin(c *gin.Context) {
	user := c.MustGet("user")
	if user != "admin" {
		response.ErrorStrResp(c, "You are not an admin", 403)
		c.Abort()
	} else {
		c.Next()
	}
}
