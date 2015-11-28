package logger

import "fmt"

// 	l.Output.Write([]byte(s))
//	s := fmt.Sprintf("[%s] : [%s] [%s::%s:%s] - %s", lvl, date, file, funct, strconv.Itoa(line), message)

func (l *Logger) messagesHandler() {
	for {
		select {
		case m := <-l.messages:
			//Here we'll have to format the file
			fmt.Println(m.format)
		}
	}
}
