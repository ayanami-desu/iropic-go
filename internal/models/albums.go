package models

import "time"

type Albums struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"unique;type:string;size:20"`
	IsR18     bool   `gorm:"default:false"`
	Desc      string `gorm:"type:string;size:200"`
	CreatedAt time.Time
	Image     []Images `gorm:"foreignKey:BelongAlbum;references:Id;constraint:OnDelete:SET NULL;"`
}
