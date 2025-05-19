package utils

import (
	"log"
	"os"
)

// Logger 定义日志接口
type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

// defaultLogger 默认日志实现
type defaultLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
}

// NewLogger 创建新的日志实例
func NewLogger() Logger {
	return &defaultLogger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		fatalLogger: log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info 记录信息日志
func (l *defaultLogger) Info(msg string, args ...interface{}) {
	l.infoLogger.Printf(msg, args...)
}

// Error 记录错误日志
func (l *defaultLogger) Error(msg string, args ...interface{}) {
	l.errorLogger.Printf(msg, args...)
}

// Fatal 记录致命错误并终止程序
func (l *defaultLogger) Fatal(msg string, args ...interface{}) {
	l.fatalLogger.Fatalf(msg, args...)
}

// 全局日志实例
var Log = NewLogger()
