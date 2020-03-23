package log

import (
	"blog-bxd-service/config"
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"log"
	"time"
	"github.com/gin-gonic/gin"
)

func SetupLogger(env string) {
	// 生成文件
	fileName := config.Get("project.midd_log")
	// 日志文件分片
	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d%H",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(72*time.Hour),
		rotatelogs.WithRotationTime(time.Second*60),
	)
	if err != nil {
		fmt.Printf("[SetupLogger rotatelogsNew fail][%S]", fileName)
	}
	// 输出
	log.SetOutput(logWriter)
}

func LogerMiddleware(env string) gin.HandlerFunc {
	fileName := config.Get("project.midd_request_log")
	// 日志文件分片
	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d%H",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(72*time.Hour),
		rotatelogs.WithRotationTime(time.Second*60),
	)
	if err != nil {
		fmt.Printf("[SetupLogger rotatelogsNew fail][%S]", fileName)
	}

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIp := c.ClientIP()

		// 日志格式
		fmt.Fprintf(
			logWriter,
			"[status_code:%d][latency_time:%s][client_ip:%s][req_method:%s][req_uri:%s]\n",
			statusCode,
			latencyTime,
			clientIp,
			reqMethod,
			reqUri,
		)
	}
}