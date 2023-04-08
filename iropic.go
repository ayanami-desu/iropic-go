package main

import (
	"fmt"
	"github.com/ayanami-desu/iropic-go/api"
	"github.com/ayanami-desu/iropic-go/internal/auth"
	"github.com/ayanami-desu/iropic-go/internal/bootstrap"
	"github.com/ayanami-desu/iropic-go/internal/db"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	bootstrap.InitLog()
	bootstrap.InitConf()
	r := gin.Default()
	if bootstrap.Conf.IsDev == "no" {
		gin.SetMode(gin.ReleaseMode)
	}
	log.Infof("start iropic service with %s mode", gin.Mode())
	r.MaxMultipartMemory = 16 << 20
	auth.Init()
	api.Init(r)
	db.Init()
	err := r.Run(fmt.Sprintf(":%s", bootstrap.Conf.Port))
	if err != nil {
		log.Fatalf("gin启动失败%+v", err)
	}
}
