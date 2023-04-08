package db

import (
	"fmt"
	"github.com/ayanami-desu/iropic-go/internal/bootstrap"
	"github.com/ayanami-desu/iropic-go/internal/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db gorm.DB

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		bootstrap.Conf.DBConfig.User,
		bootstrap.Conf.DBConfig.Password,
		bootstrap.Conf.DBConfig.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
	if err != nil {
		log.Fatalf("failed to connect database: %+v", err)
	}
	Db = *db
	err = Db.AutoMigrate(new(models.Albums), new(models.Images), new(models.Labels))
	if err != nil {
		log.Fatalf("failed to migrate models: %+v", err)
	}
}
