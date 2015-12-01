package logger

import (
	"fmt"
	"os"
)

//Type is the one type used to define CONSOLE, FILE, ... - the type of our logger
type Type uint8

type Level uint8

type Color string

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

const (
	Debug     string = "DEBUG"
	Info      string = "INFO"
	Notice    string = "NOTICE"
	Warning   string = "WARNING"
	Error     string = "ERROR"
	Critical  string = "CRITICAL"
	Alert     string = "ALERT"
	Emergency string = "EMERGENCY"
)

const (
	DEBUG Level = iota
	INFO
	NOTICE
	WARNING
	ERROR
	CRITITAL
	ALERT
	EMERGENCY
)

const (
	Bold          string = "\033[1mbold"
	LightGrey     string = "\033[37m"
	Grey          string = "\033[39m"
	Yellow        string = "\033[33m"
	Red           string = "\033[31m"
	Green         string = "\033[32m"
	LightRed      string = "\033[91m"
	White         string = "\033[97m"
	LightBlue     string = "\033[94m"
	LightYellow   string = "\033[93m"
	BackgroundRed string = "\033[41m"
)

type loggerMessage struct {
	format string
	ltype  string
	date   string
	file   string
	funct  string
	line   int
	level  Level
}

type loggerInstance struct {
	actived Status
	output  *os.File
	ltype   Type
}

//Logger struct is the one exported. This struct is filled and returned in the Init function and will be used to stores messages and loggerInstances
type Logger struct {
	actived       bool
	messages      chan *loggerMessage
	instances     []loggerInstance
	colors        map[string]string
	colorsEnabled bool
}

//AddFileLogger open the file given as path, create a new logger and fill fields of this struct. The function returns a *Logger
func (l *Logger) AddFileLogger(path string) error {
	if _, err := os.Stat(path); os.IsExist(err) {
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

func (l *Logger) EnableColor() {
	l.colorsEnabled = true
}

func (l *Logger) DisableColor() {
	l.colorsEnabled = false
}

func (l *Logger) CheckColorStatus() bool {
	return l.colorsEnabled
}

func (l *Logger) initColorsMap() {
	l.colors[Debug] = Grey
	l.colors[Info] = White
	l.colors[Notice] = White
	l.colors[Warning] = LightYellow
	l.colors[Error] = Yellow
	l.colors[Critical] = Red
	l.colors[Alert] = Red
	l.colors[Emergency] = BackgroundRed
}

//Init method permit to init a new Logging system and return a pointer to this logger system. It will be used to add Loggers and Print messages.
func Init() *Logger {
	l := Logger{
		actived:       true,
		messages:      make(chan *loggerMessage, 64),
		colors:        make(map[string]string),
		colorsEnabled: true,
	}
	l.initColorsMap()
	go l.messagesHandler()
	return &l
}
