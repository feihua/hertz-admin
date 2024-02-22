// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"hertz_admin/biz/dal"
	"hertz_admin/biz/handler"
	"hertz_admin/biz/pkg/mw"
	"io"
	"log"
	"os"
	"path"
	"time"
)

func main() {
	if initLog() {
		return
	}

	dal.Init()
	mw.InitJwt()
	//h := server.Default()
	h := server.New(server.WithHostPorts("0.0.0.0:6666"))

	auth := h.Group("/auth", mw.JwtMiddleware.MiddlewareFunc())
	auth.GET("/ping", handler.Ping)
	register(h)
	h.Spin()
}

func initLog() bool {
	// Customizable output directory.
	var logFilePath string
	dir := "./hlog"
	logFilePath = dir + "/logs/"
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		log.Println(err.Error())
		return true
	}

	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return true
		}
	}

	encoder := zapcore.NewJSONEncoder(humanEncoderConfig())
	// For zap detailed settings, please refer to https://github.com/hertz-contrib/logger/tree/main/zap and https://github.com/uber-go/zap
	logger := hertzzap.NewLogger(hertzzap.WithCoreEnc(encoder), hertzzap.WithZapOptions(zap.AddCaller(), zap.AddCallerSkip(3)))
	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}

	fileWriter := io.MultiWriter(lumberjackLogger, os.Stdout)
	logger.SetOutput(fileWriter)
	logger.SetLevel(hlog.LevelDebug)

	hlog.SetLogger(logger)
	return false
}
func humanEncoderConfig() zapcore.EncoderConfig {
	cfg := testEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncodeDuration = zapcore.StringDurationEncoder
	return cfg
}
func testEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "name",
		TimeKey:        "ts",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
