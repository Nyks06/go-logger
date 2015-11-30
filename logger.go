package logger

import (
	"fmt"
	"os"
)

//Type is the one type used to define CONSOLE, FILE, ... - the type of our logger
type Type uint8

//Status is just a bool used to change the status of a certain type of loggers
type Status bool

const (
	//CONSOLE is the one LoggerType used in the NewConsoleLogger() function
	CONSOLE Type = iota
	//FILE is the one LoggerType used in the NewFileLogger() function
	FILE Type = iota
	//ANY is the one LoggerType used to contains CONSOLE and FILE types
	ANY Type = iota
)

type loggerMessage struct {
	format string
	ltype  string
	date   string
	file   string
	funct  string
	line   int
}

type loggerInstance struct {
	actived Status
	output  *os.File
	ltype   Type
}

type Logger struct {
	actived   bool
	messages  chan *loggerMessage
	instances []loggerInstance
}

//addFileLogger open the file given as path, create a new logger and fill fields of this struct. The function returns a *Logger
func (l *Logger) AddFileLogger(path string) error {
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
func (l *Logger) AddConsoleLogger(out *os.File) error {
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
		if elem.ltype == t {
			elem.actived = s
		}
	}
}

func (l *Logger) Enable() {
	l.actived = true
}

func (l *Logger) Disable() {
	l.actived = false
}

func (l *Logger) CheckStatus() bool {
	return l.actived
}

func (l *Logger) Quit() {
	for _, elem := range l.instances {
		if elem.ltype == FILE {
			elem.output.Close()
		}
	}
}

func Init() *Logger {
	l := Logger{
		actived:  true,
		messages: make(chan *loggerMessage, 64),
	}
	go l.messagesHandler()
	return &l
}
