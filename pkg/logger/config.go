package logger

// Config 日志配置

type Config struct {
	// 日志级别: debug, info, warn, error
	Level string `json:"level" yaml:"level"`
	// 是否输出到控制台
	Console bool `json:"console" yaml:"console"`
	// 日志文件目录
	Dir string `json:"dir" yaml:"dir"`
}

// DefaultConfig 默认配置
func DefaultConfig() *Config {
	return &Config{
		Level:   "info",
		Console: true,
		Dir:     "./logs",
	}
}
