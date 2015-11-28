package logger

import (
	"fmt"
	"os"
)

//Type is the one type used to define CONSOLE, FILE, ... - the type of our logger
type Type uint8
type Status bool

const (
	//CONSOLE is the one LoggerType used in the NewConsoleLogger() function
	CONSOLE Type = iota
	//FILE is the one LoggerType used in the NewFileLogger() function
	FILE Type = iota
	//ANY is the one LoggerType used to contains CONSOLE and FILE types
	ANY Type = iota
)

// //Logger is the main type that will be used. It contains informations about the logger as the type, a *os.file, params associated to this logger...
// type Logger struct {
// 	levels map[Level]bool
// 	Output *os.File
// 	active bool
// 	ltype  Type
// }

type loggerMessage struct {
}

type loggerInstance struct {
	actived bool
	output  *os.File
	ltype   Type
}

type Logger struct {
	actived   bool
	messages  chan *loggerMessage
	instances []loggerInstance
}

// var LevelsList = [...]Level{DEBUG, INFO, NOTICE, WARNING, ERROR, FATAL}

//addFileLogger open the file given as path, create a new logger and fill fields of this struct. The function returns a *Logger
func (l *Logger) addFileLogger(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("[GO-LOGGER] - ERROR - The file given as parameter exist - %s\n", err)
		return nil
	}
	out, err := os.Create(path)
	if err != nil {
		fmt.Printf("[GO-LOGGER] - ERROR - Can't create the file given as parameter - %s\n", err)
		return nil
	}
	i := loggerInstance{
		actived: true,
		output:  out,
		ltype:   FILE,
	}
	l.instances = append(l.instances, i)
	return nil
}

//addConsoleLogger create a Logger struct and fill fields of this struct. The function returns a *Logger
func (l *Logger) addConsoleLogger(out *os.File) error {
	i := loggerInstance{
		actived: true,
		output:  out,
		ltype:   CONSOLE,
	}
	l.instances = append(l.instances, i)
	return nil
}

func (l *Logger) ChangeStatus(t Type, s Status) {
	for _, elem := range l.instances {
		if elem.outType == t {
			elem.actived = s
		}
	}
}

func (l *Logger) Enable() {
	l.active = true
}

func (l *Logger) Disable() {
	l.active = false
}

func (l *Logger) CheckStatus() bool {
	return l.actived
}

func Init() *Logger {
	go messagesHandler()
	l := Logger{
		actived:  true,
		messages: make(chan *messages, 64),
	}
	return &l
}
