package logger

import (
	"fmt"
	"os"
)

//Type is the one type used to define CONSOLE, FILE, ... - the type of our logger
type Type uint8

const (
	//CONSOLE is the one LoggerType used in the NewConsoleLogger() function
	CONSOLE Type = iota
	//FILE is the one LoggerType used in the NewFileLogger() function
	FILE Type = iota
)

//Logger is the main type that will be used. It contains informations about the logger as the type, a *os.file, params associated to this logger...
type Logger struct {
	levels map[Level]bool
	Output *os.File
	active bool
	ltype  Type
}

//NewFileLogger open the file given as path, create a new logger and fill fields of this struct. The function returns a *Logger
func NewFileLogger(path string) *Logger {
	out, err := os.Open(path)
	if err != nil {
		fmt.Println("[GO-LOGGER] - ERROR - Can't open the file given as parameter")
		return nil
	}
	l := Logger{
		levels: make(map[Level]bool),
		Output: out,
		active: true,
		ltype:  FILE,
	}
	LevelsList := [...]Level{DEBUG, INFO, NOTICE, WARNING, ERROR, FATAL}

	for idx := range LevelsList {
		l.levels[LevelsList[idx]] = true
	}
	return &l
}

//NewConsoleLogger create a Logger struct and fill fields of this struct. The function returns a *Logger
func NewConsoleLogger(out *os.File) *Logger {
	l := Logger{
		levels: make(map[Level]bool),
		Output: out,
		active: true,
		ltype:  CONSOLE,
	}
	LevelsList := [...]Level{DEBUG, INFO, NOTICE, WARNING, ERROR, FATAL}

	for idx := range LevelsList {
		l.levels[LevelsList[idx]] = true
	}
	return &l
}

func (l *Logger) Enable() {
	l.active = true
}

func (l *Logger) Disable() {
	l.active = false
}

func (l *Logger) CheckStatus() bool {
	return l.active
}
