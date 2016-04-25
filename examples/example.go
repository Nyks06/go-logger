package main

import (
	"os"

	"github.com/nyks06/go-logger"
)

func main() {

	//First of all, you need to Init the logger.
	//You'll have to do it only one time in your program and you'll have to save the returned pointer.

	//These functions show you how to add a logger type.
	//You can add as much logger as you want, even with the same type (excepted for syslog, for the moment).
	logger.AddConsoleLogger(os.Stderr)
	logger.AddFileLogger("./log.file")
	//Windows way to add a syslog logger
	logger.AddSyslogLogger("udp", "127.0.0.1:514", "syslog-Tag")
	//Unix way to add a syslog logger
	logger.AddSyslogLogger("", "", "syslog-Tag")

	//You can enable or disable the color display with a simple function
	//By default, the text is in color. (Configurable only for console output)
	//You can check in the color display is enabled or not with a function returning a bool
	logger.DisableColor()
	logger.EnableColor()
	if logger.CheckColorStatus() {
		//Color is enabled
	} else {
		//Color is not enabled
	}

	//You also can enable or disable the logger at every moment with a simple function
	//As you can check the status of the color display, you can check the status of the logger.
	//When the logger is disabled, messages will be trashed. By default the logger is enabled.
	logger.Disable()
	logger.Enable()
	if logger.IsEnabled() {
		//Logger is enabled
	} else {
		//Logger is disabled
	}

	//You easily can configure other things like enable or disable file, console or syslog logging.
	logger.DisableConsoleLogger()
	logger.DisableFileLogger()
	logger.DisableSyslogLogger()
	logger.EnableConsoleLogger()
	logger.EnableFileLogger()
	logger.EnableSyslogLogger()

	//You can display message using functions with the name corresponding to the level of log you want.
	//You can format your log message as if you were using log package function fmt.Printf
	logger.Debug("This is a %s message", "debug")
	logger.Info("This is an info message - %d", 42)
	logger.Notice("This is a notice message")
	logger.Warning("This is a warning message")
	logger.Error("This is an error message")
	logger.Critical("This is a critical error message")
	logger.Alert("This is an alert error message")
	logger.Emergency("This is a emergency message")

	//Finally, you can close all Opened files with the logger.Quit() function
	logger.Close()
}
