package bclog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// error logger
var errorLogger *zap.SugaredLogger

func Config() {
	fileName := "log/zap.log"
	level := getLoggerLevel("debug")
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	errorLogger = logger.Sugar()
}

func Debug(args ...interface{}) {
	errorLogger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	errorLogger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	errorLogger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	errorLogger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	errorLogger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	errorLogger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	errorLogger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	errorLogger.Errorf(format, args...)
}

func DPanic(args ...interface{}) {
	errorLogger.DPanic(args...)
}

func DPanicf(format string, args ...interface{}) {
	errorLogger.DPanicf(format, args...)
}

func Panic(args ...interface{}) {
	errorLogger.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	errorLogger.Panicf(format, args...)
}

func Fatal(args ...interface{}) {
	errorLogger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	errorLogger.Fatalf(format, args...)
}
