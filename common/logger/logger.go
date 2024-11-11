package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*logrus.Logger
}

// LogConfig .
type LogConfig struct {
	Level      string // dev/info/warn/error/fatal/panic
	Filename   string // 日志文件路径
	MaxSize    int    // 每个日志文件的最大大小（MB）
	MaxBackups int    // 保留的旧日志文件的最大数量
	MaxAge     int    // 保留旧日志文件的最大天数
}

func NewLogger(conf *LogConfig) *Logger {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.SetOutput(&lumberjack.Logger{
		Filename:   conf.Filename,   // 日志文件路径
		MaxSize:    conf.MaxSize,    // 每个日志文件的最大大小（MB）
		MaxBackups: conf.MaxBackups, // 保留的旧日志文件的最大数量
		MaxAge:     conf.MaxAge,     // 保留旧日志文件的最大天数
		LocalTime:  true,            // 使用本地时间
		Compress:   true,            // 压缩旧日志文件
	})

	switch conf.Level {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	}

	return &Logger{Logger: logger}
}
