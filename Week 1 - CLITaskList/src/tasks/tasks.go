// Handles CRUD on Tasks
package tasks

/*
Handles Tasks Struct, and Task Management.
*/

import (
	"fmt"
)

type Task struct {
	id          int
	name        string
	description string
}

var Tasks []Task

func getNextTaskID() int { return len(Tasks) + 1 }

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
	return Task{id: getNextTaskID(), name: name, description: desc}
}

func AddTaskToSlice(t Task) {
	Tasks = append(Tasks, t) // Overwrite original slice, with new slice made by append with new task.
	fmt.Println("Task Added.")
}
