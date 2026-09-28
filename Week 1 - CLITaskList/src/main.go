package main

/*
Brief: Main Starting Module
Command-Line Arguments and Parsing User Input to Tasks and Storage
*/

import (
	"fmt"
	"os"
	"tasklist/tasks"
)

// Program Starting Point
func main() {
	// Ensure Arguments Present
	if len(os.Args) == 1 {
		fmt.Println("No Arguments Provided")
		return
	}

	// Ensure Subcommand Exists
	subcommand := os.Args[1] // option

	if subcommand == "all" {
		tasks.ShowAllTasks()
		return
	}

	firstArg := os.Args[2]  // task name
	secondArg := os.Args[3] // task description

	if firstArg == "" {
		fmt.Println("Missing Task Name")
		return
	}

	switch subcommand {
	case "Add":
		if secondArg == "" {
			fmt.Println("Missing Task Description")
			return
		}

	case "Remove":
		break
	}
}
