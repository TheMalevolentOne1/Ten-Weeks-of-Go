// Handles CRUD on Tasks
package tasks

/*
Handles Tasks Struct, and Task Management.
*/

import (
	"fmt"
)

type Task struct {
	name        string
	description string
}

var Tasks []Task

func ShowAllTasks() {
	if lenTasks := len(Tasks); lenTasks == 0 {
		fmt.Println("No Tasks.")
	} else {
		for i := 0; i < lenTasks; i++ {
			fmt.Println(Tasks[i])
		}
	}
}

func CreateTask(name string, desc string) Task {
	fmt.Println("Adding Task")

	newTask := Task{}
	newTask.name = name
	newTask.description = desc

	return newTask
}

func AddTaskToList(t Task) {
	// append(Tasks, t) - Something something pointer slices indexes something sigh.
	fmt.Println("Task Added.")
	return
}
