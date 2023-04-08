package handles

import (
	"github.com/ayanami-desu/iropic-go/api/response"
	"github.com/ayanami-desu/iropic-go/internal/auth"
	"github.com/ayanami-desu/iropic-go/internal/bootstrap"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type userReq struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func LoginHandle(c *gin.Context) {
	var user userReq
	if err := c.ShouldBind(&user); err != nil {
		response.ErrorResp(c, err, 400)
		return
	}
	if user.Username != bootstrap.Conf.Username || user.Password != bootstrap.Conf.Password {
		response.ErrorStrResp(c, "账号或密码不正确", 400)
		return
	}
	token, err := auth.GenerateToken("admin")
	if err != nil {
		log.Warnf("failed to generate token")
		response.ErrorResp(c, err, 500)
		return
	}
	response.SuccessResp(c, gin.H{"token": token})
}
func CheckLogin(c *gin.Context) {
	user := c.MustGet("user")
	if user == "admin" {
		response.SuccessStrResp(c, "登录状态有效")
	} else {
		response.ErrorStrResp(c, "请登录", 401)
	}
}
