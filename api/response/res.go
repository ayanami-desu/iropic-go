package response

import (
	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ErrorResp is used to return error response
func ErrorResp(c *gin.Context, err error, code int) {

	c.JSON(code, Resp{
		Code: code,
		Msg:  err.Error(),
		Data: nil,
	})
	c.Abort()
}

func ErrorStrResp(c *gin.Context, str string, code int) {

	c.JSON(code, Resp{
		Code: code,
		Msg:  str,
		Data: nil,
	})
	c.Abort()
}

func SuccessResp(c *gin.Context, data ...interface{}) {
	if len(data) == 0 {
		c.JSON(200, Resp{
			Code: 200,
			Msg:  "success",
			Data: nil,
		})
		return
	}
	c.JSON(200, Resp{
		Code: 200,
		Msg:  "success",
		Data: data[0],
	})
}
func SuccessStrResp(c *gin.Context, str string) {

	c.JSON(200, Resp{
		Code: 200,
		Msg:  str,
		Data: nil,
	})
	c.Abort()
}
