package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	colorReset = "\033[0m"

	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func greeting() {
	fmt.Println("\nExample usage:")

	fmt.Println(colorYellow + "-v, --view")
	fmt.Println(colorReset + "\tDisplays the default task list.")

	fmt.Println(colorYellow + "-n, --new")
	fmt.Println(colorReset + "\tCreates a new task.")

	fmt.Println(colorYellow + "-P, --new-project")
	fmt.Println(colorReset + "\tCreates a new project")
}

func main() {
	fmt.Println(colorBlue + "Go-Do-It CLI task list by CyanCoding")

	configFile, _ :=os.UserConfigDir()
	configFile += "/go-do-it"
	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		greeting()
	}
}
