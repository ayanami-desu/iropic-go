package models

type Images struct {
	Id           uint     `gorm:"primaryKey;autoIncrement"`
	Label        []Labels `gorm:"many2many:images_labels;constraint:OnDelete:CASCADE"`
	BelongAlbum  uint     `gorm:"default:null"`
	Md5          string   `gorm:"unique;type:string;size:50"`
	Size         string   `gorm:"unique;type:string;size:15"`
	LastModified string   `gorm:"type:string;size:20"`
	IsR18        bool     `gorm:"default:false"`
	SubImg       string   `gorm:"type:string;size:100"`
	IsSub        bool     `gorm:"default:false"`
	FileType     string   `gorm:"type:string;size:10"`
	FileName     string   `gorm:"type:string;size:200"`
	OriginFile   string   `gorm:"type:string;size:250"`
	WebpFile     string   `gorm:"type:string;size:250"`
}
