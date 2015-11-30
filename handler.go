package logger

import (
	"fmt"
	"strconv"
)

func (l *Logger) formatMessage(m *loggerMessage) string {
	log := fmt.Sprintf("[%s] : [%s] [%s::%s:%s] - %s", m.ltype, m.date, m.file, m.funct, strconv.Itoa(m.line), m.format)
	return log
}

func (l *Logger) printMessage(log string) {
	for _, i := range l.instances {
		i.output.Write([]byte(log))
	}
}

func (l *Logger) messagesHandler() {
	for {
		select {
		case m := <-l.messages:
			log := l.formatMessage(m)
			l.printMessage(log)
		}
	}
}
