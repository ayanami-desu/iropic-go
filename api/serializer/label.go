package serializer

import (
	"github.com/ayanami-desu/iropic-go/internal/db"
	"github.com/ayanami-desu/iropic-go/internal/models"
)

type LabelInImgInfo struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
type LabelItem struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	ImageNum int64  `json:"imageNum"`
}

func NewLabelList(labels []models.Labels) ([]LabelItem, error) {
	var ls []LabelItem
	for _, label := range labels {
		var num int64
		db.Db.Table("images_labels").Where("labels_id = ?", label.Id).Count(&num)
		ls = append(ls, LabelItem{
			Id:       label.Id,
			Name:     label.Name,
			ImageNum: num,
		})
	}
	return ls, nil
}
