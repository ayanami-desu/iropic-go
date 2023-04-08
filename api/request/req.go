package request

type CreateImg struct {
	FileName     string `form:"name" binding:"required"`
	Size         int64  `form:"size" binding:"required"`
	LastModified string `form:"lastModified" binding:"required"`
	FileType     string `form:"fileType" binding:"required"`
	Aid          uint   `form:"aid"`
}
type CreateAlbum struct {
	Name  string `form:"name" binding:"required"`
	Desc  string `form:"desc"`
	IsR18 bool   `form:"isR18"`
}
type ImageList struct {
	Page  int    `form:"page" binding:"required"`
	Num   int    `form:"num" binding:"required"`
	Aid   uint   `form:"aid"`
	Tags  string `form:"tags"`
	Order string `form:"order" binding:"required"`
	Mode  string `form:"mode"`
}
type ImageLabel struct {
	ImagesId string `form:"pid" binding:"required"`
	LabelsId string `form:"lid" binding:"required"`
}
type AlbumEdit struct {
	Aid   uint   `form:"aid" binding:"required"`
	Name  string `form:"name" binding:"required"`
	IsR18 bool   `form:"isR18" binding:"required"`
	Desc  string `form:"desc" binding:"required"`
}
