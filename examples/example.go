package main

import (
	"os"

	"github.com/nyks06/go-logger"
)

func main() {

	//First of all, you need to Init the logger.
	//You'll have to do it only one time in your program and you'll have to save the returned pointer.
	l := logger.Init()

	//These functions show you how to add a logger type.
	//You can add as much logger as you want, even with the same type (excepted for syslog, for the moment).
	l.AddConsoleLogger(os.Stderr)
	l.AddFileLogger("./log.file")
	//Windows way to add a syslog logger
	l.AddSyslogLogger("udp", "127.0.0.1:514", "syslog-Tag")
	//Unix way to add a syslog logger
	l.AddSyslogLogger("", "", "syslog-Tag")

	//You can enable or disable the color display with a simple function
	//By default, the text is in color. (Configurable only for console output)
	//You can check in the color display is enabled or not with a function returning a bool
	l.DisableColor()
	l.EnableColor()
	if l.CheckColorStatus() {
		//Color is enabled
	} else {
		//Color is not enabled
	}

	//You also can enable or disable the logger at every moment with a simple function
	//As you can check the status of the color display, you can check the status of the logger.
	//When the logger is disabled, messages will be trashed. By default the logger is enabled.
	l.Disable()
	l.Enable()
	if l.CheckStatus() {
		//Logger is enabled
	} else {
		//Logger is disabled
	}

	//You easily can configure other things like enable or disable file, console or syslog logging.
	l.DisableConsoleLogger()
	l.DisableFileLogger()
	l.DisableSyslogLogger()
	l.EnableConsoleLogger()
	l.EnableFileLogger()
	l.EnableSyslogLogger()

	//You can display message using functions with the name corresponding to the level of log you want.
	//You can format your log message as if you were using log package function fmt.Printf
	l.Debug("This is a %s message", "debug")
	l.Info("This is an info message - %d", 42)
	l.Notice("This is a notice message")
	l.Warning("This is a warning message")
	l.Error("This is an error message")
	l.Critical("This is a critical error message")
	l.Alert("This is an alert error message")
	l.Emergency("This is a emergency message")

	//The logger act as a singleton. You can get the instance using Get(). If you've not call Init(), it will do it and returns the result.
	l2 := logger.Get()
	l2.Emergency("Message printed using Get()")

	//Finally, you can close all Opened files with the l.Quit() function
	l.Quit()
}
