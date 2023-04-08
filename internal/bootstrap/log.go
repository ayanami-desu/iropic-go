package bootstrap

import (
	"github.com/ayanami-desu/iropic-go/internal/utils"
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLog() {
	log.SetLevel(log.InfoLevel)
	if err := utils.CheckPath("./data/log"); err != nil {
		log.Fatalf("创建日志文件夹失败%+v", err)
	}
	writer, err := os.OpenFile("./data/log/log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %+v", err)
	}
	log.SetOutput(writer)
}
