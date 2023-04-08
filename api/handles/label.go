package handles

import (
	"github.com/ayanami-desu/iropic-go/api/response"
	"github.com/ayanami-desu/iropic-go/api/serializer"
	"github.com/ayanami-desu/iropic-go/internal/db"
	"github.com/ayanami-desu/iropic-go/internal/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

func CreateLabel(c *gin.Context) {
	labelStr := c.PostForm("labelStr")
	if labelStr == "" {
		response.ErrorStrResp(c, "labelStr不能为空", 400)
		return
	}
	labelList := strings.Split(labelStr, ",")
	var labels []models.Labels
	for _, name := range labelList {
		labels = append(labels, models.Labels{Name: name})
	}
	if err := db.Db.Create(&labels).Error; err != nil {
		log.Warnf("failed to create labels, labelNameList is %s; %+v", labelStr, err)
		response.ErrorResp(c, err, 500)
		return
	}
	response.SuccessStrResp(c, "创建标签成功")

}

func DeleteLabel(c *gin.Context) {
	lid := c.Query("lid")
	if lid == "" {
		response.ErrorStrResp(c, "lid不能为空", 400)
		return
	}
	if err := db.Db.Delete(&models.Labels{}, lid).Error; err != nil {
		log.Warnf("failed to delete labels, lid is %s; %+v", lid, err)
		response.ErrorResp(c, err, 500)
		return
	}
	response.SuccessStrResp(c, "删除标签成功")
}
func GetLabelList(c *gin.Context) {
	var ls []models.Labels
	if err := db.Db.Find(&ls).Error; err != nil {
		log.Warnf("failed to find labels; %+v", err)
		response.ErrorResp(c, err, 500)
		return
	}
	data, err := serializer.NewLabelList(ls)
	if err != nil {
		response.ErrorResp(c, err, 500)
		log.Warnf("failed to serialize labels; %+v", err)
		return
	}
	response.SuccessResp(c, data)
}
func SearchLabel(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		response.ErrorStrResp(c, "key不能为空", 400)
		return
	}
	var labels []models.Labels
	if err := db.Db.Where("name LIKE ?", "%"+key+"%").Find(&labels).Error; err != nil {
		log.Warnf("failed to find labels with keyword %s; %+v", key, err)
		response.ErrorResp(c, err, 500)
		return
	}
	data, err := serializer.NewLabelList(labels)
	if err != nil {
		log.Warnf("failed to serialize labels; %+v", err)
		response.ErrorResp(c, err, 500)
		return
	}
	response.SuccessResp(c, data)
}
