package main

/*
Brief: Main Starting Module

Main - Command-Line Arguments ONLY.
Tasks - Handles the Struct, and Slice.
Storage - Handles File Handling and JSON Un/Marshaling
*/

import (
	"fmt"            // Format Text To Display in Terminal
	"os"             // File Handling
	"strconv"        // Verify and Convert String Arg to ID Int - Variable Type
	"strings"        // String Manipulation Functions
	"tasklist/tasks" // Tasks Exported Functions
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

	if !tasks.EnsureStorageExists() {
		return
	}

	tasks.DecodeAllJSONTasks() // Decode JSON Data and Populate Tasks Struct.

	if subcommand == "all" {
		tasks.ShowAllTasks()
		return
	}

	firstArg := os.Args[2]

	if len(os.Args[2]) > 0 && firstArg == "" {
		fmt.Println("Missing Task Name")
		return
	}

	switch subcommand {
	case "add":
		var newTask = tasks.Task{}
		if taskName := os.Args[3]; taskName != "" {
			newTask = tasks.Task{Id: tasks.GetNextTaskID(), Name: firstArg, Description: taskName}
		} else {
			newTask = tasks.Task{Id: tasks.GetNextTaskID(), Name: firstArg, Description: taskName}
		}
		tasks.SaveNewTask(newTask)
		fmt.Println("Task Saved.")
	case "Remove":
		if len(os.Args) <= 2 {
			fmt.Println("Missing Task ID to Remove.")
			return
		}

		taskId := os.Args[2]

		if id, err := strconv.Atoi(taskId); err != nil {
			fmt.Println("That is not a number.")
		} else {
			if tPos, err := tasks.FindIDTask(id); err == nil {
				tasks.RemoveTask(tPos)
			} else {
				fmt.Println(err)
			}
		}
	}
}
