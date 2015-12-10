package logger

import (
	"fmt"
	"os"
	"strconv"
)

func (l *Logger) checkIfTTY(i *loggerInstance) bool {
	stat, _ := i.output.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return false
	}
	return true
}

func (l *Logger) shouldDisplayColor(i *loggerInstance) bool {
	if i.ltype == CONSOLE && l.colorsEnabled == true && l.checkIfTTY(i) == true {
		return true
	}
	return false
}

func (l *Logger) syslogPrintMessage(ltype string, log string) {
	if l.syslog.enabled == false {
		return
	}
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

func (l *Logger) formatMessage(m *loggerMessage) (string, string) {
	log := fmt.Sprintf("[%s] : [%s] [%s::%s:%s] - %s\n", m.ltype, m.date, m.file, m.funct, strconv.Itoa(m.line), m.format)
	logMin := fmt.Sprintf("[%s::%s:%s] - %s\n", m.file, m.funct, strconv.Itoa(m.line), m.format)
	return log, logMin
}

func (l *Logger) printMessage(m *loggerMessage, log string, logMin string) {
	for _, i := range l.instances {
		if i.enabled == true {
			if l.shouldDisplayColor(&i) {
				i.output.Write([]byte(l.colors[m.ltype]))
			}
			i.output.Write([]byte(log))
			if l.shouldDisplayColor(&i) {
				i.output.Write([]byte(Reset))
			}
		}
	}
	l.syslogPrintMessage(m.ltype, logMin)
}

func (l *Logger) handledMessage(m *loggerMessage) {
	log, logMin := l.formatMessage(m)
	if l.enabled == true {
		l.printMessage(m, log, logMin)
	}
}
