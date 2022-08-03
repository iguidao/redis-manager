package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/iguidao/redis-web-manager/src/cfg"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// error logger
var ErrorLogger *zap.SugaredLogger

func SetupLogger() *zap.SugaredLogger {
	err := os.Mkdir("./logs/", os.ModePerm)
	if err != nil && !strings.Contains(err.Error(), "exists") {
		fmt.Println("create logs dir err, ", err.Error())
	}
	fileName := cfg.Get_Local("logapppath")
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   1024,
		LocalTime: true,
		Compress:  true,
	})
	encoder := zap.NewDevelopmentEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	var level zapcore.Level
	level = zap.DebugLevel
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			syncWriter),
		level,
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	ErrorLogger = logger.Sugar()
	return ErrorLogger
}

func Debug(args ...interface{}) {
	ErrorLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	ErrorLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	ErrorLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	ErrorLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	ErrorLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	ErrorLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	ErrorLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	ErrorLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	ErrorLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	ErrorLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	ErrorLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	ErrorLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	ErrorLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	ErrorLogger.Fatalf(template, args...)
}
