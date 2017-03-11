package logging

import "fmt"

func PrintError(str string) {
	fmt.Println("Error:", str)
}

func PrintInfo(str string) {
	fmt.Println("Info:", str)
}

func PrintWarning(str string) {
	fmt.Println("Warning:", str)
}
