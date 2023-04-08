package handles

import (
	"github.com/ayanami-desu/iropic-go/api/response"
	"github.com/ayanami-desu/iropic-go/api/serializer"
	"github.com/ayanami-desu/iropic-go/internal/db"
	"github.com/ayanami-desu/iropic-go/internal/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAlbumList(c *gin.Context) {
	var albums []models.Albums
	if err := db.Db.Find(&albums).Error; err != nil {
		log.Warnf("failed to find album list; %+v", err)
		response.ErrorResp(c, err, 500)
		return
	}
	ls, _ := serializer.NewAlbumList(albums)
	response.SuccessResp(c, ls)
}
