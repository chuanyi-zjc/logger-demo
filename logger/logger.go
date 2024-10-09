package logger

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

type LogLevel int

const (
    DEBUG LogLevel = iota
    INFO
    WARNING
    ERROR
)

type Logger struct {
    minLevel LogLevel
    mutex    sync.Mutex
}

func NewLogger(minLevel LogLevel) *Logger {
    return &Logger{minLevel: minLevel}
}

func (l *Logger) log(level LogLevel, format string, args...interface{}) {
    if level < l.minLevel {
        return
    }
    l.mutex.Lock()
    defer l.mutex.Unlock()
    now := time.Now()
    pc, _, line, _ := runtime.Caller(2)
    funcName := runtime.FuncForPC(pc).Name()
    fmt.Printf("%s [%d-%d] %s:%d %s: ", now.Format("2006-01-02 15:04:05"), time.Now().UnixNano(), runtime.NumGoroutine(), funcName, line, getLevelName(level))
    fmt.Printf(format, args...)
    fmt.Println()
}

func (l *Logger) Debug(format string, args...interface{}) {
    l.log(DEBUG, format, args...)
}

func (l *Logger) Info(format string, args...interface{}) {
    l.log(INFO, format, args...)
}

func (l *Logger) Warning(format string, args...interface{}) {
    l.log(WARNING, format, args...)
}

func (l *Logger) Error(format string, args...interface{}) {
    l.log(ERROR, format, args...)
}

func getLevelName(level LogLevel) string {
    switch level {
    case DEBUG:
        return "DEBUG"
    case INFO:
        return "INFO"
    case WARNING:
        return "WARNING"
    case ERROR:
        return "ERROR"
    default:
        return "UNKNOWN"
    }
}
