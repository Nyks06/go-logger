package logger

import "fmt"

func (l *Logger) log(lvl string, message string) {

	s, err := formatLog(lvl, 2, message)
	if err != nil {
		panic(err)
	}
	l.Output.Write([]byte(s))
}

func (l *Logger) Debug(format string, a ...interface{}) {
	f := fmt.Sprintf(format, a...)
	l.log("DEBUG", f)
}

func (l *Logger) Info(format string, a ...interface{}) {
	f := fmt.Sprintf(format, a...)
	l.log("INFO", f)
}

func (l *Logger) Notice(format string, a ...interface{}) {
	f := fmt.Sprintf(format, a...)
	l.log("NOTICE", f)
}

func (l *Logger) Warning(format string, a ...interface{}) {
	f := fmt.Sprintf(format, a...)
	l.log("WARNING", f)
}

func (l *Logger) Error(format string, a ...interface{}) {
	f := fmt.Sprintf(format, a...)
	l.log("ERROR", f)
}

func (l *Logger) Fatal(format string, a ...interface{}) {
	f := fmt.Sprintf(format, a...)
	l.log("FATAL", f)
}
