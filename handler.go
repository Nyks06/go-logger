package Logger

import "fmt"

func (l *Logger) messagesHandler() {
	for {
		select {
		case m, ok := <-l.messages:
			if !ok {
				fmt.Println("[GO-LOGGER - ERROR - Channel of messages is closed")
			}
			fmt.Println(m.format)
		}
	}
}
