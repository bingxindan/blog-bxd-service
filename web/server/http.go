package server

import (
	"context"
	"bxd-middleware-service/config"
	bootstrapOrder "bxd-middleware-service/gateways/order/bootstrap"
	bootstrapSale "bxd-middleware-service/gateways/sale/bootstrap"
	"bxd-middleware-service/utils/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SetServer(env string) {
	// pass this context down the chain
	ctx, cancel := context.WithCancel(context.Background())

	log.Printf("starting web server at %s%s", config.Get("project.host"), config.Get("project.port"))

	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	server := &http.Server{
		Addr:    config.Get("project.port"),
		Handler: SetRouter(env),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Fatalf("web server shutdown complete: %s", err)
			} else {
				log.Fatalf("web server closed unexpect: %s", err)
			}
		}
	}()

	<-ctx.Done()
	err := server.Close()
	if err != nil {
		log.Fatalf("web server shutdown failed: %v", err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Print("shutting down web server success")
	cancel()

	time.Sleep(3 * time.Second)
}

func SetRouter(env string) http.Handler {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	app := gin.New()

	app.Use(log.LogerMiddleware(env), gin.Recovery())

	// 接入所有项目初始化文件
	bootstrapSale.StartSaleServer(app)
	bootstrapOrder.StartOrderServer(app)

	return app
}