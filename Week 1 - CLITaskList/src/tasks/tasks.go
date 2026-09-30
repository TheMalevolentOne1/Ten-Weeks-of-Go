// Handles CRUD on Tasks
package tasks

/*
Handles Tasks Struct, and Task Management.
*/

import (
	"fmt"
)

type Task struct {
	Id          int
	Name        string
	Description string
}

var Tasks []Task

func GetNextTaskID() int { return len(Tasks) + 1 }

func ShowAllTasks() {
	if lenTasks := len(Tasks); lenTasks == 0 {
		fmt.Println("No Tasks.")
	} else {
		for i := 0; i < lenTasks; i++ {
			fmt.Println(Tasks[i])
		}
	}
}

func AddTaskToSlice(t Task) {
	Tasks = append(Tasks, t) // Overwrite original slice, with new slice made by append with new task.
	fmt.Println("Task Added.")
}
