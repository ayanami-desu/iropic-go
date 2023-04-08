package handles

import (
	"github.com/ayanami-desu/iropic-go/api/request"
	"github.com/ayanami-desu/iropic-go/api/response"
	"github.com/ayanami-desu/iropic-go/api/serializer"
	"github.com/ayanami-desu/iropic-go/internal/db"
	"github.com/ayanami-desu/iropic-go/internal/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

func GetImageInfo(c *gin.Context) {
	pid := c.Query("pid")
	if pid == "" {
		response.ErrorStrResp(c, "pid不能为空", 400)
		return
	}
	var i models.Images
	if err := db.Db.Preload("Label").First(&i, pid).Error; err != nil {
		log.Warnf("failed to find image info, pid is %s; %+v", pid, err)
		response.ErrorResp(c, err, 500)
		return
	}
	if i.IsR18 {
		user := c.MustGet("user")
		if user != "admin" {
			response.ErrorStrResp(c, "you don't have access to this source", 401)
			return
		}
	}
	info, err := serializer.NewImgInfo(&i)
	if err != nil {
		log.Warnf("failed to serialize image info, pid is %s; %+v", pid, err)
		response.ErrorResp(c, err, 500)
		return
	}
	response.SuccessResp(c, info)
}
func AddLabelToImage(c *gin.Context) {
	pid := c.PostForm("pid")
	if pid == "" {
		response.ErrorStrResp(c, "pid为空", 400)
		return
	}
	lidStr := c.PostForm("lidList")
	if lidStr == "" {
		response.ErrorStrResp(c, "lidList为空", 400)
		return
	}
	lidList := strings.Split(lidStr, ",")
	var dataMap []map[string]interface{}
	for _, lid := range lidList {
		dataMap = append(dataMap, map[string]interface{}{"images_id": pid, "labels_id": lid})
	}
	if err := db.Db.Table("images_labels").
		Create(dataMap).Error; err != nil {
		log.Warnf("failed to create many2many, pid is %s, lidList is %s; %+v", pid, lidStr, err)
		response.ErrorResp(c, err, 500)
		return
	}
	response.SuccessStrResp(c, "添加成功")
}
func DelImageLabel(c *gin.Context) {
	var form request.ImageLabel
	if err := c.ShouldBind(&form); err != nil {
		response.ErrorResp(c, err, 400)
		return
	}
	if err := db.Db.Table("images_labels").
		Where("images_id=? AND labels_id=?", form.ImagesId, form.LabelsId).
		Delete(&form).Error; err != nil {
		log.Warnf("failed to delete image's label, pid is %s, lid is %s; %+v", form.ImagesId, form.LabelsId, err)
		response.ErrorResp(c, err, 500)
		return
	}
	response.SuccessStrResp(c, "删除成功")
}
func SetImgR18(c *gin.Context) {
	pidStr := c.PostForm("pidList")
	if pidStr == "" {
		response.ErrorStrResp(c, "pidList为空", 400)
		return
	}
	pidList := strings.Split(pidStr, ",")
	if err := db.Db.Model(&models.Images{}).Where("id IN ?", pidList).Update("is_r18", 1).Error; err != nil {
		log.Warnf("failed to set image isR18, pidList is %s; %+v", pidStr, err)
		response.ErrorResp(c, err, 500)
		return
	}
	response.SuccessStrResp(c, "设置隐藏成功")
}
func CancelImgR18(c *gin.Context) {
	pidStr := c.PostForm("pidList")
	if pidStr == "" {
		response.ErrorStrResp(c, "pidList为空", 400)
		return
	}
	pidList := strings.Split(pidStr, ",")
	if err := db.Db.Model(&models.Images{}).Where("id IN ?", pidList).Update("is_r18", 0).Error; err != nil {
		log.Warnf("failed to cancel image isR18, pidList is %s; %+v", pidStr, err)
		response.ErrorResp(c, err, 500)
		return
	}
	response.SuccessStrResp(c, "取消隐藏成功")
}
