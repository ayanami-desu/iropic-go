package handles

import (
	"github.com/ayanami-desu/iropic-go/api/request"
	"github.com/ayanami-desu/iropic-go/api/response"
	"github.com/ayanami-desu/iropic-go/api/serializer"
	"github.com/ayanami-desu/iropic-go/internal/bootstrap"
	"github.com/ayanami-desu/iropic-go/internal/db"
	"github.com/ayanami-desu/iropic-go/internal/models"
	"github.com/ayanami-desu/iropic-go/internal/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func GetImageList(c *gin.Context) {
	user := c.MustGet("user")
	var form request.ImageList
	if err := c.Bind(&form); err != nil {
		response.ErrorResp(c, err, 400)
		return
	}
	if form.Num > bootstrap.Conf.PageMaxNum || form.Page < 1 || form.Num < 1 {
		response.ErrorStrResp(c, "page或num不合法", 400)
		return
	}
	var images []models.Images
	_images := db.Db.Table("images").Where("is_sub=?", 0)
	if user != "admin" {
		_images = _images.Where("is_r18=?", 0)
	}
	if form.Aid != 0 {
		_images = _images.Where("belong_album=?", form.Aid)
	}
	if form.Tags != "" {
		TagList := strings.Split(form.Tags, ",")
		ids, _ := getPidList(TagList[0])
		for i := 1; i < len(TagList); i++ {
			ls, _ := getPidList(TagList[i])
			//错误处理
			ids = utils.Intersect(ids, ls)
		}
		_images = _images.Where("id IN ?", ids)

		//以上是并模式，还需要或模式
	}

	var imageNum int64
	_images.Count(&imageNum)

	if err := _images.Limit(form.Num).
		Offset((form.Page - 1) * form.Num).
		Order(form.Order).Preload("Label").
		Find(&images).Error; err != nil {
		log.Warnf("failed to find images; %+v", err)
		response.ErrorResp(c, err, 500)
		return
	}
	ls, err := serializer.NewImgList(images)
	if err != nil {
		log.Warnf("failed to serialize images; %+v", err)
		response.ErrorResp(c, err, 500)
	}
	response.SuccessResp(c, gin.H{
		"data":     ls,
		"imageNum": imageNum,
	})
}

func GetImgWithoutLabel(c *gin.Context) {
	var images []models.Images
	err := db.Db.Joins("left outer join images_labels on images.id = images_labels.images_id").
		Where("images_labels.labels_id IS NULL").Where("is_sub=?", 0).Find(&images).Error
	if err != nil {
		log.Warnf("failed to find images without labels; %+v", err)
		response.ErrorResp(c, err, 500)
		return
	}
	ls, _ := serializer.NewImageSeq(images)
	response.SuccessResp(c, ls)
}
func RandomImage(c *gin.Context) {
	numStr := c.Query("num")
	if numStr == "" {
		response.ErrorStrResp(c, "num为空", 400)
		return
	}
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		response.ErrorResp(c, err, 400)
		return
	}
	var pidList []uint
	var imageNum int64
	db.Db.Table("images").Count(&imageNum)
	if imageNum < num {
		response.ErrorStrResp(c, "要求的图片数过多", 400)
		return
	}
	err = db.Db.Table("images").Select("id").Find(&pidList).Error
	if err != nil {
		log.Warnf("failed to find pids when load random images; %+v", err)
		response.ErrorResp(c, err, 500)
		return
	}
	ls := utils.RandomSlice(int(num), pidList)
	response.SuccessResp(c, ls)

}
func getPidList(tag string) ([]uint, error) {
	var ids []uint
	err := db.Db.Table("images_labels").
		Select("images_id").
		Where("labels_id=?", tag).
		Order("images_id").
		Find(&ids).Error
	if err != nil {
		return ids, err
	}
	return ids, nil
}
