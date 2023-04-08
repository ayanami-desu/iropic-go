package serializer

import (
	"github.com/ayanami-desu/iropic-go/internal/db"
	"github.com/ayanami-desu/iropic-go/internal/models"
)

type AlbumList struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	IsR18    bool   `json:"isR18"`
	Desc     string `json:"desc"`
	ImageNum int64  `json:"imageNum"`
}

func NewAlbumList(albums []models.Albums) ([]AlbumList, error) {
	var ls []AlbumList
	for _, album := range albums {
		num := db.Db.Model(&album).Association("Image").Count()
		ls = append(ls, AlbumList{
			Id:       album.Id,
			Name:     album.Name,
			IsR18:    album.IsR18,
			Desc:     album.Desc,
			ImageNum: num,
		})
	}
	return ls, nil
}
