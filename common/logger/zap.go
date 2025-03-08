package logger

import (
	"asset_management_backend/app_config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapSugaredLogger *zap.SugaredLogger

// func GetLogger() *zap.SugaredLogger {
// 	if zapSugaredLogger == nil {
// 		panic("Logger not initialized. Call InitAppLoggerByConfig first.")
// 	}
// 	return zapSugaredLogger
// }

func InitAppLogger() error {
	loggerCfg := app_config.GetAppConfig().LoggerConfig
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(loggerCfg.Level)); err != nil {
		return err
	}

	config := zap.Config{
		Development:       loggerCfg.Development,
		DisableCaller:     loggerCfg.DisableCaller,
		DisableStacktrace: loggerCfg.DisableStacktrace,
		Encoding:          loggerCfg.Encoding,
		Level:             zap.NewAtomicLevelAt(level),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		EncoderConfig:     zap.NewProductionEncoderConfig(),
	}
	zapLogger, err := config.Build()
	zapSugaredLogger = zapLogger.Sugar()
	return err
}

func log(level zapcore.Level, msg string, fields ...interface{}) {
	zapSugaredLogger.With(fields...).Log(level, msg)
}

func Debug(msg string, fields ...interface{}) {
	log(zapcore.DebugLevel, msg, fields...)
}

func Info(msg string, fields ...interface{}) {
	log(zapcore.InfoLevel, msg, fields...)
}

func Warn(msg string, fields ...interface{}) {
	log(zapcore.WarnLevel, msg, fields...)
}

func Error(msg string, fields ...interface{}) {
	log(zapcore.ErrorLevel, msg, fields...)
}

func Fatal(msg string, fields ...interface{}) {
	log(zapcore.FatalLevel, msg, fields...)
}

func Sync() {
	zapSugaredLogger.Sync()
}
