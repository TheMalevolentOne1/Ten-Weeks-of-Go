// Handles CRUD on Tasks
package tasks

import (
	"fmt"
	"tasklist/storage"
)

type task struct {
	name        string
	description string
}

func AddTask(name string, desc string) {
	fmt.Println("Adding Task")

	data, err := storage.Read()

	if err == nil {
		fmt.Println(string(data))
		return
	}
}
