package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gookit/color"
)

var theLogger = map[string]interface{}{
	"Error":   broadCastLog,
	"Warning": noLog,
	"Info":    noLog,
	"Debug":   noLog,
}

var broadCaster = map[string]interface{}{
	"logToConsole": noLog,
}

// InitLogger Initialize Logger to be used in allover the service
func InitLogger(LogLevel, LogConsole string) {
	if strings.ToLower(LogConsole) == "yes" {
		broadCaster["logToConsole"] = logConsole
	}
	switch LogLevel {
	case "1", "Error":
		theLogger["Debug"] = noLog
		theLogger["Info"] = noLog
		theLogger["Warning"] = noLog
		theLogger["Error"] = broadCastLog
		break
	case "2", "Warning":
		theLogger["Debug"] = noLog
		theLogger["Info"] = noLog
		theLogger["Warning"] = broadCastLog
		theLogger["Error"] = broadCastLog
		break
	case "3", "Info":
		theLogger["Debug"] = noLog
		theLogger["Info"] = broadCastLog
		theLogger["Warning"] = broadCastLog
		theLogger["Error"] = broadCastLog
		break
	default:
		theLogger["Debug"] = broadCastLog
		theLogger["Info"] = broadCastLog
		theLogger["Warning"] = broadCastLog
		theLogger["Error"] = broadCastLog
	}
	Log("Logger Initialized")
}

// Log to send logs to system-wide configured log outputs. call as Log(Data String)
func Log(Data string) {
	TxtColor := color.FgGreen.Render
	theLogger["Info"].(func(string))(TxtColor(Data))
}

// Error send logs to system-wide configured log outputs. call as Error(Data String)
func Error(Data string) {
	TxtColor := color.FgRed.Render
	theLogger["Error"].(func(string))(TxtColor(Data))
}

// Warning send logs to system-wide configured log outputs. call as Warning(Data String)
func Warning(Data string) {
	TxtColor := color.FgYellow.Render
	theLogger["Warning"].(func(string))(TxtColor(Data))
}

// Info send logs to system-wide configured log outputs. call as Info(Data String)
func Info(Data string) {
	TxtColor := color.FgBlue.Render
	theLogger["Info"].(func(string))(TxtColor(Data))
}

// Debug send logs to system-wide configured log outputs. call as Debug(Data String)
func Debug(Data string) {
	TxtColor := color.FgWhite.Render
	theLogger["Debug"].(func(string))(TxtColor(Data))
}

// Panic send logs to system-wide configured log outputs and exit. call as Panic(Data String)
func Panic(Data string) {
	TxtColor := color.FgLightRed.Render
	theLogger["Error"].(func(string))(TxtColor(Data))
	os.Exit(-1)
}

func broadCastLog(Data string) {
	broadCaster["logToConsole"].(func(string))(Data)
}
func logConsole(Data string) {
	fmt.Printf("\n%s:\t%s\n", time.Now().String(), Data)
}

func noLog(Data string) {
	return
}
