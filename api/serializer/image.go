package serializer

import (
	"github.com/ayanami-desu/iropic-go/internal/models"
	"github.com/ayanami-desu/iropic-go/internal/query"
)

type ImageInfo struct {
	Id           uint             `json:"id" `
	Label        []LabelInImgInfo `json:"label"`
	BelongAlbum  string           `json:"belongAlbum"`
	FileName     string           `json:"filename"`
	Size         string           `json:"size"`
	LastModified string           `json:"lastModified"`
	SubImg       string           `json:"subImg"`
}
type ImageList struct {
	Id           uint             `json:"id"`
	IsR18        bool             `json:"isR18"`
	SubImg       string           `json:"subImg"`
	Label        []LabelInImgInfo `json:"label"`
	LastModified string           `json:"lastModified"`
	BelongAlbum  string           `json:"belongAlbum"`
	FileName     string           `json:"filename"`
}
type ImageSeq struct {
	Id       uint   `json:"id"`
	FileName string `json:"filename"`
}

func NewImgInfo(image *models.Images) (*ImageInfo, error) {
	ba := ""
	if image.BelongAlbum != 0 {
		album, err := query.GetAlbumById(image.BelongAlbum)
		if err != nil {
			return nil, err
		}
		ba = album.Name
	}
	var labelList []LabelInImgInfo
	if len(image.Label) != 0 {
		for _, label := range image.Label {
			labelList = append(labelList, LabelInImgInfo{Id: label.Id, Name: label.Name})
		}
	}
	info := ImageInfo{
		Id:           image.Id,
		Label:        labelList,
		BelongAlbum:  ba,
		FileName:     image.FileName,
		Size:         image.Size,
		LastModified: image.LastModified,
		SubImg:       image.SubImg,
	}
	return &info, nil
}
func NewImgList(images []models.Images) ([]ImageList, error) {
	var ImageLs []ImageList
	ba := ""
	for _, image := range images {
		ba = ""
		if image.BelongAlbum != 0 {
			album, err := query.GetAlbumById(image.BelongAlbum)
			if err != nil {
				return nil, err
			}
			ba = album.Name
		}
		var labelList []LabelInImgInfo
		if len(image.Label) != 0 {
			for _, label := range image.Label {
				labelList = append(labelList, LabelInImgInfo{Id: label.Id, Name: label.Name})
			}
		}
		ImageLs = append(ImageLs, ImageList{
			Id:           image.Id,
			Label:        labelList,
			BelongAlbum:  ba,
			LastModified: image.LastModified,
			IsR18:        image.IsR18,
			SubImg:       image.SubImg,
			FileName:     image.FileName,
		})
	}
	return ImageLs, nil
}
func NewImageSeq(images []models.Images) ([]ImageSeq, error) {
	var ls []ImageSeq
	for _, image := range images {
		ls = append(ls, ImageSeq{
			Id:       image.Id,
			FileName: image.FileName,
		})
	}
	return ls, nil
}
