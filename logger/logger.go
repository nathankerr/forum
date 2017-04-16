package logger

import (
	"fmt"
	"time"
)

func Error(format string, a ...interface{}) {
	print("Error", format, a...)
}

func Info(format string, a ...interface{}) {
	print("Info", format, a...)
}

func Warning(format string, a ...interface{}) {
	print("Warning", format, a...)
}

func Debug(format string, a ...interface{}) {
	print("Debug", format, a...)
}

func print(category string, format string, a ...interface{}) {
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
	}

	fmt.Printf("%s%s - %s%s: ", colour, now, category, reset)
	fmt.Printf(format, a...)
	fmt.Printf("\n")
}
