package infras

import (
	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"runtime"
	"saas/pkg/consts"
	"time"
)

func InitLogger() {
	logFilePath := consts.KlogFilePath
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}

	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}

	logger := kitexlogrus.NewLogger()
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "",
		MaxSize:    0,
		MaxAge:     0,
		MaxBackups: 0,
		LocalTime:  false,
		Compress:   false,
	}
	if runtime.GOOS == "linux" {
		logger.SetOutput(lumberjackLogger)
		logger.SetLevel(klog.LevelWarn)
	} else {
		logger.SetLevel(klog.LevelDebug)
	}
	klog.SetLogger(logger)
}
