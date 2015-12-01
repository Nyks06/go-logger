package logger

import (
	"fmt"
	"os"
	"strconv"
)

func (l *Logger) checkIfTTY() bool {
	stat, _ := os.Stdout.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return false
	}
	return true
}

func (l *Logger) shouldDisplayColor(i *loggerInstance) bool {
	if i.ltype == CONSOLE && l.colorsEnabled == true && l.checkIfTTY() == true {
		return true
	}
	return false
}

func (l *Logger) syslogPrintMessage(ltype string, log string) {
	switch ltype {
	case Debug:
		l.syslog.Writer.Debug(log)
	case Info:
		l.syslog.Writer.Info(log)
	case Notice:
		l.syslog.Writer.Notice(log)
	case Warning:
		l.syslog.Writer.Warning(log)
	case Error:
		l.syslog.Writer.Err(log)
	case Critical:
		l.syslog.Writer.Crit(log)
	case Alert:
		l.syslog.Writer.Alert(log)
	case Emergency:
		l.syslog.Writer.Emerg(log)
	}
}

func (l *Logger) formatMessage(m *loggerMessage) string {
	log := fmt.Sprintf("[%s] : [%s] [%s::%s:%s] - %s\n", m.ltype, m.date, m.file, m.funct, strconv.Itoa(m.line), m.format)
	return log
}

func (l *Logger) printMessage(m *loggerMessage, log string) {
	for _, i := range l.instances {
		if l.shouldDisplayColor(&i) {
			i.output.Write([]byte(l.colors[m.ltype]))
		}
		i.output.Write([]byte(log))
		if l.shouldDisplayColor(&i) {
			i.output.Write([]byte(Reset))
		}
	}
	l.syslogPrintMessage(m.ltype, log)
}

func (l *Logger) messagesHandler() {
	for {
		select {
		case m := <-l.messages:
			log := l.formatMessage(m)
			l.printMessage(m, log)
		}
	}
}
