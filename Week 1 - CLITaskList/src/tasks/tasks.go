// Handles CRUD on Tasks
package tasks

import (
	"fmt"
)

type Task struct {
	name        string
	description string
}

var Tasks []Task

func ShowAllTasks() {
	for i := 0; i < len(Tasks); i++ {
		fmt.Println(Tasks[i])
	}
	return
}

func CreateTask(name string, desc string) Task {
	fmt.Println("Adding Task")

	newTask := Task{}
	newTask.name = name
	newTask.description = desc
}

func AddTaskToList(t Task) {
	// append(Tasks, t) - Something something pointer slices indexes something sigh.
	fmt.Println("Task Added.")
	return
}

// Helper Function
// Brief: Return Next Tasks Slices Index
func nextTaskIndex() int {
	if lenTasks := len(Tasks); lenTasks == 0 {
		return 0
	} else {
		return lenTasks + 1
	}
}
