package main

/*
Brief: Main Starting Module

To maintain a separation of concerns:

Main - Command-Line Arguments ONLY.
Tasks - Handles the Struct, and Slice.
Storage - Handles File Handling and JSON Un/Marshaling
*/

import (
	"fmt"
	"os"
	"strings"
	"tasklist/tasks"
)

func toLower(x string) string {
	var newStr string = ""
	for i := 0; i < len(x); i++ {
		newStr += strings.ToLower(string(x[i])) // convert x byte into string?... isn't this already a string by the parameter -_-#
	}
	return newStr
}

// Program Starting Point
func main() {
	// Ensure Arguments Present
	if len(os.Args) == 1 {
		fmt.Println("No Arguments Provided")
		return
	}

	// Ensure Subcommand Exists
	subcommand := os.Args[1] // all, add, remove

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
	case "add":
		if secondArg == "" {
			fmt.Println("Missing Task Description")
			return
		}

		newTask := tasks.CreateTask(firstArg, secondArg)
		tasks.AddTaskToSlice(newTask)

		fmt.Println("TASKS ARE NOT CURRENTLY SAVED! FUNCTIONALLY ADDED TO TASKS LIST FOR LATER REFERENCE BY STORAGE MODULE!")
	case "Remove":
		fmt.Println("TODO: REMOVE TASK")
	}
}
