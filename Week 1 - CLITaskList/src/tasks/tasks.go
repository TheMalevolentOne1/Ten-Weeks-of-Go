// Handles CRUD on Tasks
package tasks

import (
	"fmt"
	"tasklist/storage"
)

type Task struct {
	name        string
	description string
}

func CreateTask(name string, desc string) Task {
	fmt.Println("Adding Task")

	res := storage.Write(name, desc)

	if res {
		fmt.Println("Successfully write to File")
	} else {
		fmt.Println("Failure to write to File")
	}
}
