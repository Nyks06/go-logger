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

//Debug is the one method called to display a DEBUG message
func (l *Logger) Debug(format string, a ...interface{}) {
	l.generateMessage(1, "DEBUG", fmt.Sprintf(format, a...))
}

//Info is the one method called to display a INFO message
func (l *Logger) Info(format string, a ...interface{}) {
	l.generateMessage(1, "INFO", fmt.Sprintf(format, a...))
}

//Notice is the one method called to display a NOTICE message
func (l *Logger) Notice(format string, a ...interface{}) {
	l.generateMessage(1, "NOTICE", fmt.Sprintf(format, a...))
}

//Warning is the one method called to display a WARNING message
func (l *Logger) Warning(format string, a ...interface{}) {
	l.generateMessage(1, "WARNING", fmt.Sprintf(format, a...))
}

//Error is the one method called to display a ERROR message
func (l *Logger) Error(format string, a ...interface{}) {
	l.generateMessage(1, "ERROR", fmt.Sprintf(format, a...))
}

//Critical is the one method called to display a CRITICAL message
func (l *Logger) Critical(format string, a ...interface{}) {
	l.generateMessage(1, "CRITICAL", fmt.Sprintf(format, a...))
}

//Alert is the one method called to display a ALERT message
func (l *Logger) Alert(format string, a ...interface{}) {
	l.generateMessage(1, "ALERT", fmt.Sprintf(format, a...))
}

//Emergency is the one method called to display a EMERGENCY message
func (l *Logger) Emergency(format string, a ...interface{}) {
	l.generateMessage(1, "EMERGENCY", fmt.Sprintf(format, a...))
}
