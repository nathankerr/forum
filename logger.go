package main

import (
	"fmt"
	"time"
)

type colouredLogger struct{}

var logger colouredLogger

func (l colouredLogger) Error(format string, a ...interface{}) {
	l.print("Error", format, a...)
}

func (l colouredLogger) Info(format string, a ...interface{}) {
	l.print("Info", format, a...)
}

func (l colouredLogger) Warning(format string, a ...interface{}) {
	l.print("Warning", format, a...)
}

func (l colouredLogger) Debug(format string, a ...interface{}) {
	l.print("Debug", format, a...)
}

func (l colouredLogger) API(format string, a ...interface{}) {
	l.print("API", format, a...)
}

func (l colouredLogger) print(category string, format string, a ...interface{}) {
	var colour string
	reset := "\x1b[0m"
	now := time.Now().Format("2006-01-02 15:04:05")

	switch category {
	case "Error":
		colour = "\x1b[31m"
	case "Info":
		colour = "\x1b[32m"
	case "Warning":
		colour = "\x1b[33m"
	case "Debug":
		colour = "\x1b[34m"
	case "API":
		colour = "\x1b[35m"
	}

	fmt.Printf("%s%s - %s%s: ", colour, now, category, reset)
	fmt.Printf(format, a...)
	fmt.Printf("\n")
}
