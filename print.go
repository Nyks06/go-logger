package logger

import (
	"fmt"
	"time"
)

func (l *Logger) generateMessage(depth int, ltype string, format string) {
	file, funct, line, err := getInfos(depth + 1)
	if err != nil {
		panic(err)
	}

	m := loggerMessage{
		format: format,
		ltype:  ltype,
		date:   time.Now().Format("2006-01-02 15:04:05"),
		file:   file,
		funct:  funct,
		line:   line,
	}
	l.messages <- &m
}

func (l *Logger) Debug(format string, a ...interface{}) {
	l.generateMessage(1, "DEBUG", fmt.Sprintf(format, a...))
}

func (l *Logger) Info(format string, a ...interface{}) {
	l.generateMessage(1, "INFO", fmt.Sprintf(format, a...))
}

func (l *Logger) Notice(format string, a ...interface{}) {
	l.generateMessage(1, "NOTICE", fmt.Sprintf(format, a...))
}

func (l *Logger) Warning(format string, a ...interface{}) {
	l.generateMessage(1, "WARNING", fmt.Sprintf(format, a...))
}

func (l *Logger) Error(format string, a ...interface{}) {
	l.generateMessage(1, "ERROR", fmt.Sprintf(format, a...))
}

func (l *Logger) Fatal(format string, a ...interface{}) {
	l.generateMessage(1, "FATAL", fmt.Sprintf(format, a...))
}
