package logger

import (
	"fmt"
	"strconv"
)

func (l *Logger) formatMessage(m *loggerMessage) string {
	log := fmt.Sprintf("[%s] : [%s] [%s::%s:%s] - %s\n", m.ltype, m.date, m.file, m.funct, strconv.Itoa(m.line), m.format)
	return log
}

func (l *Logger) printMessage(m *loggerMessage, log string) {
	for _, i := range l.instances {
		if i.ltype == CONSOLE && l.colorsEnabled == true {
			i.output.Write([]byte(l.colors[m.ltype]))
		}
		i.output.Write([]byte(log))
		if i.ltype == CONSOLE && l.colorsEnabled == true {
			i.output.Write([]byte("\033[00m"))
		}
	}
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
