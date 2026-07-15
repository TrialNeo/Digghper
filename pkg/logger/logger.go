package logger

import (
	"Diggpher/global"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
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
	// 控制台输出
	consoleWriter := zapcore.Lock(os.Stdout)

	// 文件输出（带轮转）
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  "logs/app/app.log",
		MaxSize:   100,
		LocalTime: true,
		Compress:  false,
	})

	// 如果希望控制台也带颜色，保留 CapitalColorLevelEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 文件一般用 JSON 或不带颜色的 ConsoleEncoder（可选）
	// 这里为了简单，文件也用同一种 encoder（但去掉颜色更稳妥）
	fileEncoderConfig := encoderConfig
	fileEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 去掉颜色，避免日志文件含 ANSI 转义符
	fileEncoder := zapcore.NewConsoleEncoder(fileEncoderConfig)

	// 创建两个 core
	consoleCore := zapcore.NewCore(consoleEncoder, consoleWriter, level)
	fileCore := zapcore.NewCore(fileEncoder, fileWriter, level)

	// 合并 cores
	core := zapcore.NewTee(consoleCore, fileCore)

	global.Log = zap.New(core, zap.AddCaller())
	global.SugarLog = global.Log.Sugar()
}

// Config 日志配置结构
type Config struct {
	Level   string
	Console bool
	Dir     string
}

// DefaultConfig 返回默认日志配置
func DefaultConfig() *Config {
	return &Config{
		Level:   "info",
		Console: true,
		Dir:     "./logs",
	}
}
