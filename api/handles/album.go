package handles

import (
	"github.com/ayanami-desu/iropic-go/api/request"
	"github.com/ayanami-desu/iropic-go/api/response"
	"github.com/ayanami-desu/iropic-go/internal/bootstrap"
	"github.com/ayanami-desu/iropic-go/internal/db"
	"github.com/ayanami-desu/iropic-go/internal/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

func CreateAlbum(c *gin.Context) {
	var form request.CreateAlbum
	if err := c.Bind(&form); err != nil {
		response.ErrorResp(c, err, 400)
		return
	}
	album := models.Albums{
		Desc:  form.Desc,
		Name:  form.Name,
		IsR18: form.IsR18,
	}
	if err := db.Db.Create(&album).Error; err != nil {
		log.Warnf("failed to create album; %+v", err)
		response.ErrorResp(c, err, 400)
		return
	}
	response.SuccessStrResp(c, "创建相册成功")
}
func EditAlbum(c *gin.Context) {
	var form request.AlbumEdit
	if err := c.ShouldBind(&form); err != nil {
		response.ErrorResp(c, err, 400)
		return
	}
	var album models.Albums
	if err := db.Db.First(&album, form.Aid).Error; err != nil {
		log.Warnf("failed to find album by aid %d; %+v", form.Aid, err)
		response.ErrorResp(c, err, 500)
		return
	}
	err := db.Db.Model(&album).Updates(models.Albums{Desc: form.Desc, IsR18: form.IsR18, Name: form.Name}).Error
	if err != nil {
		log.Warnf("failed to edit album, aid is %d; %+v", form.Aid, err)
		response.ErrorResp(c, err, 500)
		return
	}
	response.SuccessStrResp(c, "更新相册成功")
}
func DelAlbum(c *gin.Context) {
	aid := c.Query("aid")
	if aid == "" {
		response.ErrorStrResp(c, "aid为空", 400)
		return
	}
	if err := db.Db.Delete(&models.Albums{}, aid).Error; err != nil {
		log.Warnf("failed to delete album, aid is %s; %+v", aid, err)
		response.ErrorResp(c, err, 500)
		return
	}
	response.SuccessStrResp(c, "删除相册成功")
}
func GetAlbumCover(c *gin.Context) {
	aid := c.Query("aid")
	if aid == "" {
		response.ErrorStrResp(c, "aid为空", 400)
		return
	}
	var img models.Images
	if err := db.Db.Where("belong_album=? AND is_r18=0", aid).First(&img).Error; err != nil {
		//log.Warnf("failed to find album cover, aid is %s; %+v", aid, err)
		c.Redirect(302, "https://http.cat/404")
		return
	}
	if img.WebpFile == "" {
		if img.OriginFile == "" {
			response.ErrorStrResp(c, "找不到图片文件", 404)
			return
		}
		c.File(bootstrap.Conf.PathPrefix + img.OriginFile)
		return
	}
	c.File(bootstrap.Conf.PathPrefix + img.WebpFile)
}
func MvImgToAlbum(c *gin.Context) {
	aid := c.PostForm("aid")
	if aid == "" {
		response.ErrorStrResp(c, "aid为空", 400)
		return
	}
	pidStr := c.PostForm("pidList")
	if pidStr == "" {
		response.ErrorStrResp(c, "pidList为空", 400)
		return
	}
	var album models.Albums
	var img models.Images
	if err := db.Db.First(&album, aid).Error; err != nil {
		log.Warnf("failed to find album, aid is %s; %+v", aid, err)
		response.ErrorResp(c, err, 500)
		return
	}
	pidList := strings.Split(pidStr, ",")
	if err := db.Db.Model(&img).Where("id IN ?", pidList).Update("belong_album", aid).Error; err != nil {
		log.Warnf("failed to move images to album, aid is %s, pidList is %s; %+v", aid, pidStr, err)
		response.ErrorResp(c, err, 500)
		return
	}
	response.SuccessStrResp(c, "成功将图片移入相册")
}
func RmAlbumImg(c *gin.Context) {
	pidStr := c.PostForm("pidList")
	if pidStr == "" {
		response.ErrorStrResp(c, "pidList为空", 400)
		return
	}
	pidList := strings.Split(pidStr, ",")
	var i models.Images
	if err := db.Db.Model(&i).Where("id IN ?", pidList).Update("belong_album", "null").Error; err != nil {
		log.Warnf("failed to remove images from album, pidList is %s; %+v", pidStr, err)
		response.ErrorResp(c, err, 500)
		return
	}
	response.SuccessStrResp(c, "成功从相册移除图片")
}
