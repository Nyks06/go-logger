package logger

func (l *Logger) log(lvl string, message string) {

	s, err := formatLog(lvl, 2, message)
	if err != nil {
		return
	}
	l.Output.Write([]byte(s))
}

func (l *Logger) Debug(message string) {
	l.log("DEBUG", message)
}

func (l *Logger) Info(message string) {
	l.log("INFO", message)
}

func (l *Logger) Notice(message string) {
	l.log("NOTICE", message)
}

func (l *Logger) Warning(message string) {
	l.log("WARNING", message)
}

func (l *Logger) Error(message string) {
	l.log("ERROR", message)
}

func (l *Logger) Fatal(message string) {
	l.log("FATAL", message)
}
