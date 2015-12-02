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
	//You can add as much logger as you want, even with the same type.
	l.AddConsoleLogger(os.Stderr)
	l.AddFileLogger("./log.file")
	l.AddSyslogLogger("syslog-Prefix")

	//You can enable or disable the color display with a simple function
	//By default, the text is in color. (Configurable only for console output)
	//You can check in the color display is enabled or not with a function returning a bool
	// l.DisableColor()
	// l.EnableColor()
	// if l.CheckColorStatus() {
	// 	//Color is enabled
	// } else {
	// 	//Color is not enabled
	// }

	//You also can enable or disable the logger at every moment with a simple function
	//As you can check the status of the color display, you can check the status of the logger.
	//When the logger is disabled, messages will be trashed. By default the logger is enabled.
	// l.Disable()
	// l.Enable()
	// if l.CheckStatus() {
	// 	//Logger is enabled
	// } else {
	// 	//Logger is disabled
	// }

	//You easily can configure other things like enable or disable file, console or syslog logging.
	// l.DisableConsoleLogger()
	// l.DisableFileLogger()
	// l.DisableSyslogLogger()
	// l.EnableConsoleLogger()
	// l.EnableFileLogger()
	// l.EnableSyslogLogger()

	//Finally, you can display message using functions with the name corresponding to the level of log you want.
	//You can format your log message as if you were using log package function fmt.Printf
	l.Debug("This is a %s message", "debug")
	l.Info("This is an info message - %d", 42)
	l.Notice("This is a notice message")
	l.Warning("This is a warning message")
	l.Error("This is an error message")
	l.Critical("This is a critical error message")
	l.Alert("This is an alert error message")
	l.Emergency("This is a emergency message")

	//Finally, you can close all Opened files with the l.Quit() function
	l.Quit()
}
