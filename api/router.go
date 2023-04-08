package api

import (
	"github.com/ayanami-desu/iropic-go/api/handles"
	"github.com/ayanami-desu/iropic-go/internal/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.Use()
	Cors(r)
	r.POST("/login", handles.LoginHandle)
	auth := r.Group("")
	auth.Use(middlewares.Auth)
	auth.GET("/test", handles.TestAuth)

	//image相关Api
	image := auth.Group("/image")
	image.GET("/info", handles.GetImageInfo)
	image.GET("/image", handles.GetImage)
	image.POST("/images", handles.GetImageList)
	image.GET("/random", handles.RandomImage)

	//album相关Api
	album := r.Group("/album")
	album.GET("/albums", handles.GetAlbumList)
	album.GET("/cover", handles.GetAlbumCover)

	//label相关Api
	label := r.Group("/label")
	label.GET("/labels", handles.GetLabelList)
	label.GET("/query", handles.SearchLabel)

	admin := auth.Group("")
	admin.Use(middlewares.AuthAdmin)
	adminApi(admin)
}
func adminApi(r *gin.RouterGroup) {
	r.GET("/login/check", handles.CheckLogin)
	image := r.Group("/image")
	//image相关Api
	image.POST("/upload", handles.CreateImage)
	image.GET("/seq", handles.GetImgWithoutLabel)
	image.DELETE("/image", handles.DelImage)
	image.POST("/label", handles.AddLabelToImage)
	image.DELETE("/label", handles.DelImageLabel)
	image.POST("/r18", handles.SetImgR18)
	image.PUT("/r18", handles.CancelImgR18)
	image.POST("/group", handles.SetImageGroup)
	album := r.Group("/album")
	//album相关Api
	album.POST("/image", handles.MvImgToAlbum)
	//album.PUT("/image", handles.RmAlbumImg)
	album.POST("/album", handles.CreateAlbum)
	album.DELETE("/album", handles.DelAlbum)
	album.PUT("/album", handles.EditAlbum)
	label := r.Group("/label")
	//label相关Api
	label.POST("/label", handles.CreateLabel)
	label.DELETE("/label", handles.DeleteLabel)

}
func Cors(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	r.Use(cors.New(config))
}
