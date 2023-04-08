package handles

import (
	"github.com/ayanami-desu/iropic-go/api/response"
	"github.com/ayanami-desu/iropic-go/internal/db"
	"github.com/ayanami-desu/iropic-go/internal/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

func SetImageGroup(c *gin.Context) {
	pidStr := c.PostForm("pidList")
	pidList := strings.Split(pidStr, ",")
	if len(pidList) < 2 {
		response.ErrorStrResp(c, "pid数量不足", 400)
		return
	}
	subImgStr := strings.Join(pidList[1:], ",")
	if err := db.Db.Where("id = ?", pidList[0]).Update("sub_img", subImgStr).Error; err != nil {
		log.Warnf("failed to set images group, pidList is %s; %+v", pidStr, err)
		response.ErrorResp(c, err, 500)
		return
	}

	if err := db.Db.Model(&models.Images{}).Where("id IN ?", pidList[1:]).Update("is_sub", 1).Error; err != nil {
		log.Warnf("failed to set image is_sub, pidList is %s; %+v", pidStr, err)
		response.ErrorResp(c, err, 500)
		return
	}

	response.SuccessStrResp(c, "成功将图片设为一组")

}
