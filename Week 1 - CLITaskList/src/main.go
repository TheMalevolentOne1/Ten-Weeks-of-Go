package main

/*
Brief: Main Starting Module

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
		newStr += strings.ToLower(string(x[i]))
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

	firstArg := os.Args[2] // task name

	if firstArg == "" {
		fmt.Println("Missing Task Name")
		return
	}

	switch subcommand {
	case "add":
		var newTask tasks.Task

		if secondArg := os.Args[3]; secondArg != "" {
			newTask = tasks.CreateTask(firstArg, secondArg)
		} else {
			newTask = tasks.CreateTask(firstArg, "")
		}

		tasks.AddTaskToSlice(newTask)
		fmt.Println("TASKS ARE NOT CURRENTLY SAVED! FUNCTIONALLY ADDED TO TASKS LIST FOR LATER REFERENCE BY STORAGE MODULE!")
	case "Remove":
		fmt.Println("TODO: REMOVE TASK")
	}
}
