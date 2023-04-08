package bootstrap

import (
	"fmt"
	"github.com/ayanami-desu/iropic-go/internal/utils"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

type DBConfig struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	DBName   string `json:"db_name,omitempty"`
}
type Config struct {
	IsDev      string   `json:"isDev,omitempty"`
	Username   string   `json:"username,omitempty"`
	Password   string   `json:"password,omitempty"`
	JwtSecret  string   `json:"jwtSecret,omitempty"`
	PageMaxNum int      `json:"pageMaxNum,omitempty"`
	PathPrefix string   `json:"pathPrefix,omitempty"`
	Port       string   `json:"port,omitempty"`
	DBConfig   DBConfig `json:"db_config,omitempty"`
}

var Conf Config

func InitConf() {

	if err := utils.CheckPath("./data"); err != nil {
		log.Fatalf("创建配置文件夹失败: %+v", err)
	}
	jsonFile, err := os.Open("./data/config.json")
	defer jsonFile.Close()
	conf := DefaultConfig()
	if err != nil {
		fmt.Printf("配置文件不存在，使用默认配置\n")
		_, err := os.Create("./data/config.json")
		if err != nil {
			log.Fatalf("创建配置文件失败%+v", err)
		}
		Conf = *conf
	} else {
		configBytes, err := io.ReadAll(jsonFile)
		if err != nil {
			log.Fatalf("读取配置文件失败%+v", err)

		}
		err = utils.Json.Unmarshal(configBytes, conf)
		if err != nil {
			log.Fatalf("更新配置文件失败%+v", err)
		}
		Conf = *conf
	}
	if !utils.WriteJsonToFile("./data/config.json", conf) {
		log.Fatalf("写入配置文件失败")
	}
	log.Infof("成功写入配置文件")

}
