package logger

import (
	"errors"
	"os"
	"sync"

	slog "github.com/Nyks06/go-syslog"
)

//Type is the one type used to define CONSOLE, FILE, ... - the type of our logger
type outType uint8

type logLevel uint8

type logColor string

//Status is just a bool used to change the status of a certain type of loggers
type status bool

const (
	consoleType outType = iota
	fileType    outType = iota
	syslogType  outType = iota
)

const (
	debugLevel     string = "DEBUG"
	infoLevel      string = "INFO"
	noticeLevel    string = "NOTICE"
	warningLevel   string = "WARNING"
	errorLevel     string = "ERROR"
	criticalLevel  string = "CRITICAL"
	alertLevel     string = "ALERT"
	emergencyLevel string = "EMERGENCY"
)

const (
	bold          string = "\033[1mbold"
	reset         string = "\033[00m"
	lightGrey     string = "\033[37m"
	grey          string = "\033[39m"
	yellow        string = "\033[33m"
	red           string = "\033[31m"
	green         string = "\033[32m"
	lightRed      string = "\033[91m"
	white         string = "\033[97m"
	lightBlue     string = "\033[94m"
	lightYellow   string = "\033[93m"
	backgroundRed string = "\033[41m"
)

type loggerMessage struct {
	format string
	ltype  string
	date   string
	file   string
	fnct   string
	line   int
	level  logLevel
}

type loggerInstance struct {
	enabled status
	output  *os.File
	ltype   outType
}

type loggerSyslog struct {
	writer  *slog.Writer
	enabled status
}

//Logger struct is the one exported. This struct is filled and returned in the Init function and will be used to stores messages and loggerInstances
type Logger struct {
	enabled status
	mutex   *sync.RWMutex

	instances []loggerInstance
	syslog    *loggerSyslog

	colors        map[string]string
	colorsEnabled bool
}

var l *Logger

func init() {
	l := new(Logger)
	l.enabled = true
	l.mutex = &sync.RWMutex{}

	l.syslog = &loggerSyslog{enabled: false}

	l.colors = make(map[string]string)
	l.colorsEnabled = true
	initColorsMap()
}

func initColorsMap() {
	l.colors[debugLevel] = grey
	l.colors[infoLevel] = white
	l.colors[noticeLevel] = white
	l.colors[warningLevel] = lightYellow
	l.colors[errorLevel] = yellow
	l.colors[criticalLevel] = red
	l.colors[alertLevel] = red
	l.colors[emergencyLevel] = backgroundRed
}

//AddFileLogger open the file given as path, create a new logger and fill fields of this struct. The function returns a *Logger
func AddFileLogger(path string) error {
	if _, err := os.Stat(path); os.IsExist(err) {
		return errors.New("[GO-LOGGER] - ERROR - The file given as parameter exist")
	}
	out, err := os.Create(path)
	if err != nil {
		return errors.New("[GO-LOGGER] - ERROR - Can't create the file given as parameter")
	}
	i := loggerInstance{
		enabled: true,
		output:  out,
		ltype:   fileType,
	}
	l.instances = append(l.instances, i)
	return nil
}

//AddConsoleLogger create a Logger struct and fill fields of this struct. The function returns a *Logger
func AddConsoleLogger(out *os.File) error {
	i := loggerInstance{
		enabled: true,
		output:  out,
		ltype:   consoleType,
	}
	l.instances = append(l.instances, i)
	return nil
}

//AddSyslogLogger function needs the Nyks06/go-syslog package to works because of the problems between Windows platform and go-syslog package.
func AddSyslogLogger(network, raddr, prefix string) error {
	s, err := slog.Dial(network, raddr, slog.LOG_DEBUG, prefix)
	if err != nil {
		return errors.New("[GO-LOGGER] - ERROR - Can't connect to syslog")
	}
	l.syslog.writer = s
	l.syslog.enabled = true
	return nil
}

// ##########################
// ####### STATUS MANAGEMENT
// ##########################

//EnableConsoleLogger function permits to enable Console loggers created before.
func EnableConsoleLogger() {
	for _, elem := range l.instances {
		if elem.ltype == consoleType {
			elem.enabled = true
		}
	}
}

//DisableConsoleLogger function permits to disable Console loggers created before.
func DisableConsoleLogger() {
	for _, elem := range l.instances {
		if elem.ltype == consoleType {
			elem.enabled = false
		}
	}
}

//EnableFileLogger function permits to enable File loggers created before.
func EnableFileLogger() {
	for _, elem := range l.instances {
		if elem.ltype == fileType {
			elem.enabled = true
		}
	}
}

//DisableFileLogger function permits to disable File loggers created before.
func DisableFileLogger() {
	for _, elem := range l.instances {
		if elem.ltype == fileType {
			elem.enabled = false
		}
	}
}

//EnableSyslogLogger function permits to enable Syslog logger if already initialized.
func EnableSyslogLogger() {
	l.syslog.enabled = true
}

//DisableSyslogLogger function permits to disable Syslog logger if already initialized.
func DisableSyslogLogger() {
	l.syslog.enabled = false
}

//Enable method permit to enable the logger system. When enabled, the logger system will print the messages when received
func Enable() {
	l.enabled = true
}

//Disable method permits to disable the logger system. When disabled, the logger system will not print anything
func Disable() {
	l.enabled = false
}

//IsEnabled method permits to check if the logger system is enable or disabled.
func IsEnabled() bool {
	return bool(l.enabled)
}

// ##########################
// ####### COLOR MANAGEMENT
// ##########################

//EnableColor function enables the color for console loggers
func EnableColor() {
	l.colorsEnabled = true
}

//DisableColor function disables the color for console loggers
func DisableColor() {
	l.colorsEnabled = false
}

//ChangeColor function needs a level already existing
func ChangeColor(lvl string, color string) {
	if _, ok := l.colors[lvl]; !ok {
		return
	}
	l.colors[lvl] = color
}

//CheckColorStatus function permits to know if the color system is enabled for console loggers.
func CheckColorStatus() bool {
	return bool(l.colorsEnabled)
}

// ##########################
// ####### INIT & QUIT
// ##########################

//Close method permit to close all fles opened for logging.
func Close() {
	for _, elem := range l.instances {
		if elem.ltype == fileType {
			elem.output.Close()
		}
	}
}
