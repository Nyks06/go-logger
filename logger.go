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

//Logger struct is the one exported. This struct is filled and returned in the Init function and will be used to stores messages and loggerInstances
type Logger struct {
	actived   bool
	messages  chan *loggerMessage
	instances []loggerInstance
}

//AddFileLogger open the file given as path, create a new logger and fill fields of this struct. The function returns a *Logger
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

//AddConsoleLogger create a Logger struct and fill fields of this struct. The function returns a *Logger
func (l *Logger) AddConsoleLogger(out *os.File) error {
	i := loggerInstance{
		actived: true,
		output:  out,
		ltype:   CONSOLE,
	}
	l.instances = append(l.instances, i)
	return nil
}

//ChangeStatus is the one utilitary method usable to change the Status (Enabled / Disabled) for a kind of instances (console/file)
func (l *Logger) ChangeStatus(t Type, s Status) {
	for _, elem := range l.instances {
		if elem.ltype == t {
			elem.actived = s
		}
	}
}

//Enable method permit to enable the logger system. When enabled, the logger system will print the messages when received
func (l *Logger) Enable() {
	l.actived = true
}

//Disable method permit to disable the logger system. When disabled, the logger system will not print anything
func (l *Logger) Disable() {
	l.actived = false
}

//CheckStatus method permit to check if the logger system is enable or disabled.
func (l *Logger) CheckStatus() bool {
	return l.actived
}

//Quit method permit to close all fles opened for logging.
func (l *Logger) Quit() {
	for _, elem := range l.instances {
		if elem.ltype == FILE {
			elem.output.Close()
		}
	}
}

//Init method permit to init a new Logging system and return a pointer to this logger system. It will be used to add Loggers and Print messages.
func Init() *Logger {
	l := Logger{
		actived:  true,
		messages: make(chan *loggerMessage, 64),
	}
	go l.messagesHandler()
	return &l
}
