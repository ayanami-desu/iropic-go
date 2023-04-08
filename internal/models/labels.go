package models

type Labels struct {
	Id   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"unique;type:string;;size:20"`
}
