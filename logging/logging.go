package logging

import "fmt"

func PrintError(message string) {
	Print("Error", message)
}

func PrintInfo(message string) {
	Print("Info", message)
}

func PrintWarning(message string) {
	Print("Warning", message)
}

func Print(category string, message string) {
	switch category {
	case "Error":
		category = "\x1b[31m" + category + "\x1b[0m"
	case "Info":
		category = "\x1b[32m" + category + "\x1b[0m"
	case "Warning":
		category = "\x1b[33m" + category + "\x1b[0m"
	}

	fmt.Printf("%s: ", category)
	fmt.Println(message)
}
