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
	if i.ltype == consoleType && l.colorsEnabled == true && l.checkIfTTY(i) == true {
		return true
	}
	return false
}

func (l *Logger) syslogPrintMessage(ltype string, log string) {
	if l.syslog.enabled == false {
		return
	}
	switch ltype {
	case debugLevel:
		l.syslog.writer.Debug(log)
	case infoLevel:
		l.syslog.writer.Info(log)
	case noticeLevel:
		l.syslog.writer.Notice(log)
	case warningLevel:
		l.syslog.writer.Warning(log)
	case errorLevel:
		l.syslog.writer.Err(log)
	case criticalLevel:
		l.syslog.writer.Crit(log)
	case alertLevel:
		l.syslog.writer.Alert(log)
	case emergencyLevel:
		l.syslog.writer.Emerg(log)
	}
}

func (l *Logger) formatMessage(m *loggerMessage) (string, string) {
	log := fmt.Sprintf("[%s] : [%s] [%s::%s:%s] - %s\n", m.ltype, m.date, m.file, m.fnct, strconv.Itoa(m.line), m.format)
	logMin := fmt.Sprintf("[%s::%s:%s] - %s\n", m.file, m.fnct, strconv.Itoa(m.line), m.format)
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
				i.output.Write([]byte(reset))
			}
		}
	}
	l.syslogPrintMessage(m.ltype, logMin)
}

func (l *Logger) handledMessage(m *loggerMessage) {
	log, logMin := l.formatMessage(m)
	if l.enabled == true {
		l.mutex.Lock()
		l.printMessage(m, log, logMin)
		l.mutex.Unlock()
	}
}
