package main

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
	subcommand := os.Args[1]

	switch subcommand {
	case "add":
		tasks.AddTask("Test", "Test")
	case "remove":
		break
	case "list":
		break
	default:
		fmt.Println("etc")
	}
}
