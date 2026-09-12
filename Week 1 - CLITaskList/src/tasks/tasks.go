// Handles CRUD on Tasks
package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type task struct {
	name        string
	description string
}

func AddTask(name string, desc string) {
	fmt.Println("Adding Task")
	path := filepath.Join(os.TempDir(), "tasks.json")

	// Verify the path to File/Directory does not have Stats
	// and does Error i.e. Does **Not** Exist.
	if _, err := os.Stat("data/tasks.json"); err != nil {
		fmt.Println("Creating Data Directory")
		os.Mkdir("data", 0755) // create data folder, with 0755 perms (Owner: rwx, Everyone Else: r-x)

		// create a tasks.json file in data directory
		// file contents stored in an empty byte array
		os.WriteFile("data/tasks.json", []byte{}, 0755)
	} else {
		dataStored, err := os.ReadFile(path)

		var tasks []task

		// if dataStored cannot parse/marshal the JSON file data against the Task Struct then it will Error.
		data := json.Unmarshal(dataStored, &tasks)

		if err == nil {
			fmt.Println(data)
		} else if data != nil {
			fmt.Println(data)
		}
	}
}
