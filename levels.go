package logger

type Level uint8

const (
	DEBUG   Level = iota
	INFO    Level = iota
	NOTICE  Level = iota
	WARNING Level = iota
	ERROR   Level = iota
	FATAL   Level = iota
)
