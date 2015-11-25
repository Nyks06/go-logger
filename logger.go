package main

import (
	"fmt"
	"os"
)

//LoggerType is the one type used to define CONSOLE, FILE, ... - the type of our logger
type LoggerType uint8

const (
	//CONSOLE is the one LoggerType used in the NewConsoleLogger() function
	CONSOLE LoggerType = iota
	//FILE is the one LoggerType used in the NewFileLogger() function
	FILE LoggerType = iota
)

//Logger is the main type that will be used. It contains informations about the logger as the type, a *os.file, params associated to this logger...
type Logger struct {
	Output *os.File
	Type   LoggerType
}

//NewFileLogger open the file given as path, create a new logger and fill fields of this struct. The function returns a *Logger
func NewFileLogger(path string) *Logger {
	l := Logger{}
	var err error

	l.Type = FILE
	l.Output, err = os.Open(path)
	if err != nil {
		fmt.Println("[GO-LOGGER] - ERROR - Can't open the file given as parameter")
		return nil
	}
	return &l
}

//NewConsoleLogger create a Logger struct and fill fields of this struct. The function returns a *Logger
func NewConsoleLogger(out *os.File) *Logger {
	l := Logger{}

	l.Type = CONSOLE
	l.Output = out
	return &l
}
