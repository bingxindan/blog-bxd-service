package config

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"log"
	"os"
	"strings"
)

var (
	Config *ini.File
)

func SetupConfig(env string) {

	switch env {
	case "prod":
	case "gray":
		gin.SetMode(gin.ReleaseMode)
		break
	case "test":
		gin.SetMode(gin.TestMode)
		break
	default:
		gin.SetMode(gin.DebugMode)
	}

	log.Print("env:", env)

	var err error

	Config, err = ini.Load(
		"/Users/lauren/work/goweb/blog-bxd-service/server.ini",
		"conf/"+env+".ini")
	if err != nil {
		log.Panic(err)
		return
	}

	// 加入环境变量
	Config.ValueMapper = os.ExpandEnv
}

func Get(key string) string {
	parts := strings.Split(key, ".")
	section := parts[0]
	keyStr := parts[1]
	// 判断是否数组，返回不同结果
	return Config.Section(section).Key(keyStr).String()
}
