package handles

import "github.com/gin-gonic/gin"

func TestAuth(c *gin.Context) {
	user := c.MustGet("user")
	if user == "admin" {
		c.JSON(200, "admin")
		return
	}
	if user == "guest" {
		c.JSON(200, "guest")
		return
	}
	c.JSON(401, "unknown user")
}
func TestPost(c *gin.Context) {
	a := c.PostForm("a")
	c.JSON(200, a)
}
