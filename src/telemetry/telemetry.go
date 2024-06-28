package telemetry

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	// lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.SugaredLogger
var AccessLogger *zap.SugaredLogger

func InitLog(level string) {

	var err error
	var logLevel zapcore.Level
	var hostname, _ = os.Hostname()

	switch level {
	case "DEBUG":
		logLevel = zap.DebugLevel
	case "INFO":
		logLevel = zap.InfoLevel
	case "WARNING":
		logLevel = zap.WarnLevel
	case "ERROR":
		logLevel = zap.ErrorLevel
	default:
		logLevel = zap.InfoLevel // fallback
	}

	zapLevel := zap.NewAtomicLevelAt(logLevel)
	encoder := zap.NewProductionEncoderConfig()
	encoder.TimeKey = "timestamp"
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig = encoder
	zapConfig.Level = zapLevel
	zapConfig.Development = false
	zapConfig.Encoding = "json"
	zapConfig.InitialFields = map[string]interface{}{"release": "105", "ProcessID": os.Getpid(), "Hostname": hostname}
	zapConfig.OutputPaths = []string{"stdout", "c:\\Open311Logs\\open311-runtime.log"}
	zapConfig.ErrorOutputPaths = []string{"stdout"}

	logger, err := zapConfig.Build()

	if err != nil {
		panic(err)
	}

	Logger = logger.Sugar()

}

func InitAccessLog() {

	// this gets called from router.go

	encoder := zap.NewProductionEncoderConfig()
	encoder.TimeKey = "timestamp"
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig = encoder
	zapConfig.Level = zap.NewAtomicLevel()

	logger, err := zapConfig.Build()

	if err != nil {
		panic(err)
	}

	AccessLogger = logger.Sugar()

}

// TODO: Apache Common Log Format:
// 127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326
