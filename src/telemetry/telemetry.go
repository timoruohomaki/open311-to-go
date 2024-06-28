package telemetry

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	// lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.SugaredLogger

func InitLog(level string) {

	var err error
	var logLevel zapcore.Level

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

	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig = encoder
	zapConfig.Level = zapLevel
	zapConfig.Development = false
	zapConfig.Encoding = "json"
	zapConfig.InitialFields = map[string]interface{}{"version": "104"}
	zapConfig.OutputPaths = []string{"stdout", "c:\\Open311Logs\\open311.log"}
	zapConfig.ErrorOutputPaths = []string{"stdout", "c:\\Open311Logs\\open311_ERROR.log"}

	logger, err := zapConfig.Build()

	if err != nil {
		panic(err)
	}

	Logger = logger.Sugar()

}
