package logging

import (
	"fmt"
	"os"
	"time"
)

// Log level
const (
	Debug = iota
	Info
	Warning
	Error
	Fatal
)

const timeFormat = "2006-01-02T15:04:05"

// Logger struct
type Logger struct {
	Name  string
	IsUTC bool
	buf   []byte
}

// New returns the logger object with the logger name
func New(name string) *Logger {
	logger := &Logger{}
	logger.Name = name
	logger.IsUTC = true
	return logger
}

// Log write the log with the specified log level
func (logger *Logger) Log(level int, levelString, msg string) {
	t := time.Now()
	if logger.IsUTC == true {
		t = t.UTC()
	}
	logString := fmt.Sprintf("[%s %s %s] %s\n", t.Format(timeFormat), levelString, logger.Name, msg)
	logger.buf = logger.buf[:0]
	logger.buf = append(logger.buf, logString...)
	for _, w := range Writers {
		if level >= w.Level {
			w.Writer.Write(logger.buf)
		}
	}
}

// Debug writes the log with debug level
func (logger *Logger) Debug(v ...interface{}) {
	logger.Log(Debug, "DEBUG", fmt.Sprint(v...))
}

// Debugf writes the log using fmt.Sprintf format
func (logger *Logger) Debugf(format string, v ...interface{}) {
	logger.Log(Debug, "DEBUG", fmt.Sprintf(format, v...))
}

// Info writes the log with info level
func (logger *Logger) Info(v ...interface{}) {
	logger.Log(Info, "INFO", fmt.Sprint(v...))
}

// Infof writes the log using fmt.Sprintf format
func (logger *Logger) Infof(format string, v ...interface{}) {
	logger.Log(Info, "INFO", fmt.Sprintf(format, v...))
}

// Warning writes the log with warning level
func (logger *Logger) Warning(v ...interface{}) {
	logger.Log(Warning, "WARNING", fmt.Sprint(v...))
}

// Warningf writes the log using fmt.Sprintf format
func (logger *Logger) Warningf(format string, v ...interface{}) {
	logger.Log(Warning, "WARNING", fmt.Sprintf(format, v...))
}

// Error writes the log with error level
func (logger *Logger) Error(v ...interface{}) {
	logger.Log(Error, "ERROR", fmt.Sprint(v...))
}

// Errorf writes the log using fmt.Sprintf format
func (logger *Logger) Errorf(format string, v ...interface{}) {
	logger.Log(Error, "ERROR", fmt.Sprintf(format, v...))
}

// Fatal writes the log with fatal level, and exit the program with code -1
func (logger *Logger) Fatal(v ...interface{}) {
	logger.Log(Fatal, "FATAL", fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf writes the log using fmt.Sprintf format, and exit the program with code -1
func (logger *Logger) Fatalf(format string, v ...interface{}) {
	logger.Log(Fatal, "FATAL", fmt.Sprintf(format, v...))
	os.Exit(1)
}
