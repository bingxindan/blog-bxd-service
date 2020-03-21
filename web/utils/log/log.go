package log

import (
	"bxd-middleware-service/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type (
	MyLogger struct {
		*logrus.Logger
	}
)

type ContextHook struct {
}

var (
	logger *MyLogger
)

func init() {
	logger = &MyLogger{
		logrus.New(),
	}
}

func Logger() *MyLogger {
	return logger
}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook ContextHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 6, 6)
	cnt := runtime.Callers(6, pc)
	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		if !strings.Contains(name, "github.com/sirupsen/logrus") {
			file, line := fu.FileLine(pc[i] - 1)
			entry.Data["file"] = path.Base(file) + ":" + fmt.Sprintf("%d", line)
			entry.Data["func"] = path.Base(name)
			break
		}

	}
	return nil
}

/**
 * 创建文件
 * @time 2020-=03-20
 * @author zm
 */
func getFile(logFileName string) (*os.File) {
	src, err := os.OpenFile(logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Errorf("[getFile OpenFile fail][%s][%s]", logFileName, err)
	}
	return src
}

func getHookLocal(fileName string) logrus.Hook {
	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d%H",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(72*time.Hour),
		rotatelogs.WithRotationTime(time.Second*60),
	)
	if err != nil {
		fmt.Printf("[getHookLocal rotatelogsNew fail][%s][%s]", fileName, err)
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfsHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		/*TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,*/
		//DisableColors:true,
	})

	return lfsHook
}

func setLevel(env string, logObj *logrus.Logger) {
	switch env {
	case "prod":
	case "gray":
	case "test":
		logObj.SetLevel(logrus.InfoLevel)
		break
	default:
		logObj.SetLevel(logrus.DebugLevel)
	}
}

func SetupLogging(env string) {
	// 生成文件
	fileName := config.Get("project.midd_log")
	src := getFile(fileName)

	// 日志文件分片
	hook := getHookLocal(fileName)
	// 设置分片实例
	logger.AddHook(hook)
	// 定义输出
	logger.Out = src
	// 设置日志层级
	setLevel(env, logger.Logger)

	/*// 开始时间
	startTime := time.Now()
	// 结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)

	// 日志格式
	logger.WithFields(logrus.Fields{
		"latency_time": latencyTime,
	}).Info()*/
}

func LogerMiddleware(env string) gin.HandlerFunc {
	fileName := config.Get("project.midd_request_log")
	src := getFile(fileName)

	// 实例化
	logMiddle := logrus.New()
	// 设置日志层级
	setLevel(env, logMiddle)
	// 定义输出
	logMiddle.Out = src
	// 设置分片实例
	hook := getHookLocal(fileName)
	logMiddle.AddHook(hook)

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
		logMiddle.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIp,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}

// Print logger.Print
func Print(i ...interface{}) {
	logger.Print(i...)
}

// Printf logger.Printf
func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

// Debug logger.Debug
func Debug(i ...interface{}) {
	logger.Debug(i...)
}

// Debugf logger.Debugf
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Info logger.Info
func Info(i ...interface{}) {
	logger.Info(i...)
}

// Infof logger.Infof
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warn logger.Warn
func Warn(i ...interface{}) {
	logger.Warn(i...)
}

// Warnf logger.Warnf
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Error logger.Error
func Error(i ...interface{}) {
	logger.Error(i...)
}

// Errorf logger.Errorf
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Fatal logger.Fatal
func Fatal(i ...interface{}) {
	logger.Fatal(i...)
}

// Fatalf logger.Fatalf
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Panic logger.Panic
func Panic(i ...interface{}) {
	logger.Panic(i...)
}

// Panicf logger.Panicf
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}
