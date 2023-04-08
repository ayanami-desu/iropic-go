package query

import "github.com/ayanami-desu/iropic-go/internal/models"
import "github.com/ayanami-desu/iropic-go/internal/db"

func GetAlbumById(id uint) (*models.Albums, error) {
	var album models.Albums
	if err := db.Db.First(&album, id).Error; err != nil {
		return nil, err
	}
	return &album, nil
}
