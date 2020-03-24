package database

import (
	"blog-bxd-service/config"
	"blog-bxd-service/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

var (
	GormDB *gorm.DB
)

func Instance(connKey string) {
	var err error

	// 从库添加
	if connKey == "" || len(connKey) == 0 {
		log.Fatal("[mysql:conn:fail]")
		os.Exit(1)
	}

	conns := config.Get(connKey+".connect")
	c := utils.JsonToMap(conns)

	s := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		c["username"], c["password"],
		c["hostname"], c["database"],
		c["charset"],
	)
	GormDB, err = gorm.Open("mysql", s)

	if err != nil {
		log.Fatalf("Instance GormDB err [%s]", err)
		os.Exit(2)
	}
}
