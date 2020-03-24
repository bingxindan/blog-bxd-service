package main

import (
	"blog-bxd-service/config"
	"blog-bxd-service/server"
	"blog-bxd-service/utils/database"
	"blog-bxd-service/utils/log"
	"flag"
)

var (
	env      string
	showHelp bool
)

func init() {
	flag.StringVar(&env, "env", "local", "environment for server:[local|test|dev|beta|gray|prod]")
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.Parse()
}

func main() {
	if showHelp {
		flag.PrintDefaults()
		return
	}

	// 设置配置
	config.SetupConfig(env)

	// 设置日志
	log.SetupLogger(env)

	// 设置DB
	database.Instance("mysql_blog")

	// 设置服务
	server.SetServer(env)
}
