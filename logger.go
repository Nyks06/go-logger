package logger

import (
	"fmt"
	"log/syslog"
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
	//SYSLOG is the one LoggerType used in the NewSyslogLogger() function
	ANSYSLOG Type = iota
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

const ()

const (
	Bold          string = "\033[1mbold"
	Reset         string = "\033[00m"
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
	enabled Status
	output  *os.File
	ltype   Type
}

type loggerSyslog struct {
	Writer  *syslog.Writer
	enabled Status
}

//Logger struct is the one exported. This struct is filled and returned in the Init function and will be used to stores messages and loggerInstances
type Logger struct {
	enabled       Status
	messages      chan *loggerMessage
	instances     []loggerInstance
	colors        map[string]string
	syslog        *loggerSyslog
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
		enabled: true,
		output:  out,
		ltype:   FILE,
	}
	l.instances = append(l.instances, i)
	return nil
}

//AddConsoleLogger create a Logger struct and fill fields of this struct. The function returns a *Logger
func (l *Logger) AddConsoleLogger(out *os.File) error {
	i := loggerInstance{
		enabled: true,
		output:  out,
		ltype:   CONSOLE,
	}
	l.instances = append(l.instances, i)
	return nil
}

func (l *Logger) AddSyslogLogger(prefix string) error {
	s, err := syslog.New(syslog.LOG_DEBUG, prefix)
	if err != nil {
		fmt.Printf("[GO-LOGGER] - ERROR - Can't connect to syslog - %s\n", err)
		return nil
	}
	l.syslog.Writer = s
	l.syslog.enabled = true
	return nil
}

// ##########################
// ####### STATUS MANAGEMENT
// ##########################

//ChangeStatus is the one utilitary method usable to change the Status (Enabled / Disabled) for a kind of instances (console/file)
func (l *Logger) ChangeStatus(t Type, s Status) {
	for _, elem := range l.instances {
		if elem.ltype == t {
			elem.enabled = s
		}
	}
}

func (l *Logger) EnableConsoleLogger() {
	for _, elem := range l.instances {
		if elem.ltype == CONSOLE {
			elem.enabled = true
		}
	}
}

func (l *Logger) DisableConsoleLogger() {
	for _, elem := range l.instances {
		if elem.ltype == CONSOLE {
			elem.enabled = false
		}
	}
}

func (l *Logger) EnableFileLogger() {
	for _, elem := range l.instances {
		if elem.ltype == FILE {
			elem.enabled = true
		}
	}
}

func (l *Logger) DisableFileLogger() {
	for _, elem := range l.instances {
		if elem.ltype == FILE {
			elem.enabled = false
		}
	}
}

func (l *Logger) EnableSyslogLogger() {
	l.syslog.enabled = true
}

func (l *Logger) DisableSyslogLogger() {
	l.syslog.enabled = false
}

//Enable method permit to enable the logger system. When enabled, the logger system will print the messages when received
func (l *Logger) Enable() {
	l.enabled = true
}

//Disable method permit to disable the logger system. When disabled, the logger system will not print anything
func (l *Logger) Disable() {
	l.enabled = false
}

//CheckStatus method permit to check if the logger system is enable or disabled.
func (l *Logger) CheckStatus() bool {
	if l.enabled == true {
		return true
	}
	return false
}

// ##########################
// ####### COLOR MANAGEMENT
// ##########################

func (l *Logger) EnableColor() {
	l.colorsEnabled = true
}

func (l *Logger) DisableColor() {
	l.colorsEnabled = false
}

func (l *Logger) ChangeColor(lvl string, color string) {
	l.colors[lvl] = color
}

func (l *Logger) CheckColorStatus() bool {
	if l.colorsEnabled == true {
		return true
	}
	return false
}

// ##########################
// ####### INIT & QUIT
// ##########################

//Quit method permit to close all fles opened for logging.
func (l *Logger) Quit() {
	for _, elem := range l.instances {
		if elem.ltype == FILE {
			elem.output.Close()
		}
	}
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
		enabled:       true,
		messages:      make(chan *loggerMessage, 64),
		colors:        make(map[string]string),
		colorsEnabled: true,
	}
	l.syslog = &loggerSyslog{enabled: false}

	l.initColorsMap()
	go l.messagesHandler()
	return &l
}
