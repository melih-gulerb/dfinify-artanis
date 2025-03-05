package logging

import (
	"fmt"
	"time"
)

type LogLevel int

const (
	INFO LogLevel = iota
	WARNING
	ERROR
	PANIC
)

func Log(level LogLevel, message string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	levelStr := getLogLevelString(level)

	fmt.Printf("[%s] [%s] %s\n", currentTime, levelStr, message)

	if level == PANIC {
		panic(message)
	}
}

func getLogLevelString(level LogLevel) string {
	switch level {
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case PANIC:
		return "PANIC"
	default:
		return "UNKNOWN"
	}
}
