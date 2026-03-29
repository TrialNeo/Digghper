package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var (
	// Log 全局日志实例
	Log *zap.Logger
)

// InitLogger 初始化日志
func InitLogger(config *Config) {
	// 配置编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 配置日志级别
	var level zapcore.Level
	switch config.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	// 配置输出
	var core zapcore.Core
	if config.Console {
		// 控制台输出
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			level,
		)
	} else {
		// 文件输出
		// 确保日志目录存在
		if err := os.MkdirAll(config.Dir, 0755); err != nil {
			panic(err)
		}

		// 创建日志文件
		logFile := config.Dir + "/" + time.Now().Format("2006-01-02") + ".log"
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(file),
			level,
		)
	}

	// 创建日志实例
	Log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// Debug 输出调试级别日志
func Debug(msg string, fields ...zapcore.Field) {
	Log.Debug(msg, fields...)
}

// Info 输出信息级别日志
func Info(msg string, fields ...zapcore.Field) {
	Log.Info(msg, fields...)
}

// Warn 输出警告级别日志
func Warn(msg string, fields ...zapcore.Field) {
	Log.Warn(msg, fields...)
}

// Error 输出错误级别日志
func Error(msg string, fields ...zapcore.Field) {
	Log.Error(msg, fields...)
}

// Fatal 输出致命级别日志并退出
func Fatal(msg string, fields ...zapcore.Field) {
	Log.Fatal(msg, fields...)
}
