package handles

import (
	"fmt"
	"github.com/ayanami-desu/iropic-go/api/request"
	"github.com/ayanami-desu/iropic-go/api/response"
	"github.com/ayanami-desu/iropic-go/internal/bootstrap"
	"github.com/ayanami-desu/iropic-go/internal/db"
	"github.com/ayanami-desu/iropic-go/internal/models"
	"github.com/ayanami-desu/iropic-go/internal/optimize"
	"github.com/ayanami-desu/iropic-go/internal/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetImage(c *gin.Context) {
	pid := c.Query("pid")
	if pid == "" {
		response.ErrorStrResp(c, "参数有误", 400)
		return
	}
	t := c.DefaultQuery("type", "webp")
	var i models.Images
	if err := db.Db.First(&i, pid).Error; err != nil {
		log.Warnf("failed to find image, pid is %s; %+v", pid, err)
		response.ErrorResp(c, err, 500)
		return
	}
	if i.IsR18 {
		user := c.MustGet("user")
		if user != "admin" {
			response.ErrorStrResp(c, "you don't have access to this source", 403)
			return
		}
	}
	if i.WebpFile == "" {
		if i.OriginFile == "" {
			response.ErrorStrResp(c, "找不到图片文件", 404)
			return
		}
		c.File(bootstrap.Conf.PathPrefix + i.OriginFile)
		return
	}
	if t == "origin" {
		c.File(bootstrap.Conf.PathPrefix + i.OriginFile)
		return
	}
	c.File(bootstrap.Conf.PathPrefix + i.WebpFile)
}

func CreateImage(c *gin.Context) {
	var form request.CreateImg
	if err := c.Bind(&form); err != nil {
		response.ErrorResp(c, err, 400)
		return
	}
	if !strings.HasPrefix(form.FileType, "image/") {
		response.ErrorStrResp(c, "不支持的文件类型", 402)
		return
	}
	originFile, err := c.FormFile("file")
	if err != nil {
		response.ErrorResp(c, err, 404)
		return
	}
	randomStr := utils.RandomString(6)
	lastModified, err := strconv.ParseInt(form.LastModified, 10, 64)
	if err != nil {
		log.Warnf("failed to convert string to int64, string is %s; %+v", form.LastModified, err)
		response.ErrorResp(c, err, 500)
		return
	}
	timePath := time.UnixMilli(lastModified).Format("2006/01/02/")
	fullPath := fmt.Sprintf("%s%s%s", bootstrap.Conf.PathPrefix, bootstrap.OriginImgPath, timePath)
	NewFileName := fmt.Sprintf("%s_%s", randomStr, form.FileName)
	originFilePath := fmt.Sprintf("%s%s%s", bootstrap.OriginImgPath, timePath, NewFileName)

	if err != nil {
		response.ErrorResp(c, err, 500)
		return
	}

	md5Str, err := utils.CalcMd5(originFile)
	if err != nil {
		log.Warnf("failed to calculate md5; %+v", err)
		response.ErrorResp(c, err, 500)
		return
	}

	name := strings.Split(NewFileName, ".")
	webpFilePath := ""
	if form.Size > 300<<10 {
		if err := optimize.Conv2Webp(originFile, timePath, name[0], form.FileType); err != nil {
			log.Warnf("failed to optimize image; %+v", err)
			response.ErrorResp(c, err, 500)
			return
		}
		webpFilePath = fmt.Sprintf("%s%s%s.webp", bootstrap.WebpImgPath, timePath, name[0])
	} else {
		webpFilePath = originFilePath
	}

	img := models.Images{
		FileName:     form.FileName,
		Size:         strconv.FormatInt(form.Size, 10),
		FileType:     form.FileType,
		LastModified: form.LastModified,
		OriginFile:   originFilePath,
		WebpFile:     webpFilePath,
		Md5:          md5Str,
	}
	if form.Aid != 0 {
		img.BelongAlbum = form.Aid
	}
	if err := db.Db.Create(&img).Error; err != nil {
		log.Warnf("failed to create image; %+v", err)
		response.ErrorResp(c, err, 500)
		return
	}

	if err := utils.CheckPath(fullPath); err != nil {
		log.Warnf("failed to mkdir; %+v", err)
		response.ErrorStrResp(c, "创建文件夹时出错", 500)
		return
	}
	if err := c.SaveUploadedFile(originFile, fullPath+NewFileName); err != nil {
		log.Warnf("failed to save image file; %+v", err)
		response.ErrorResp(c, err, 500)
		return
	}
	response.SuccessStrResp(c, "上传图片成功")

}
func DelImage(c *gin.Context) {
	pidStr := c.Query("pidList")
	if pidStr == "" {
		response.ErrorStrResp(c, "pidList为空", 400)
		return
	}
	pidList := strings.Split(pidStr, ",")
	for _, pid := range pidList {
		var i models.Images
		if err := db.Db.First(&i, pid).Error; err != nil {
			response.ErrorResp(c, err, 500)
			return
		}
		if err := db.Db.Delete(&i).Error; err != nil {
			response.ErrorResp(c, err, 500)
			return
		}
		//删除文件
		if i.OriginFile != "" {
			if err := os.Remove(bootstrap.Conf.PathPrefix + i.OriginFile); err != nil {
				log.Warnf("failed to delete file, path is %s; %+v", i.OriginFile, err)
			}
		}
		if i.WebpFile != "" {
			if err := os.Remove(bootstrap.Conf.PathPrefix + i.WebpFile); err != nil {
				log.Warnf("failed to delete file, path is %s; %+v", i.WebpFile, err)
			}
		}
	}
	response.SuccessStrResp(c, "删除图片完成")
}
