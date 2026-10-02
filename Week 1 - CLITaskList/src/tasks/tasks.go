// Handles CRUD on Tasks
package tasks

/*
Handles Tasks Struct, and Task Management.
*/

import (
	"encoding/json" // For JSON handling
	"fmt"           // Format for console display
	"os"            // OS for arguments
	"path/filepath" // For Filepath handling
)

type Task struct {
	Id          int
	Name        string
	Description string
}

var Tasks []Task

func GetNextTaskID() int {
	fmt.Println(Tasks)
	return len(Tasks) + 1
}

func ShowAllTasks() {
	if lenTasks := len(Tasks); lenTasks == 0 {
		fmt.Println("No Tasks.")
	} else {
		for i := 0; i < lenTasks; i++ {
			fmt.Println(Tasks[i])
		}
	}
}

const file = "data/tasks.json"

var dir, err = os.Getwd() // GetWorkingDirectory

var filePath string = filepath.Join(dir, file) // json data file path

/*
Brief: Encodes Task Struct from Struct into JSON
Parameters:
b - bytes (the bytes of the Tasks slice)
*/
func encodeJSONTask(t Task) []byte {
	data, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	} else {
		return data
	}
}

/*
Brief: Decodes all bytes of JSON into Tasks Struct
Parameters:
b - bytes (the bytes of JSON data)
*/
func DecodeAllJSONTasks() {
	if b, err := os.ReadFile(filePath); err == nil {
		err = json.Unmarshal(b, &Tasks)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func doesStorageExist() bool {
	if _, err := os.Stat(file); err == nil {
		if data, err := os.ReadFile(file); err == nil {
			// confirm length contains at least two and [ and ] and CORRECT BYTES [91, 93]
			if len(data) >= 2 && data[0] == '[' && data[len(data)-1] == ']' {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

/*
Brief: Ensure the storage exists, and create it.
Returnn Boolean whether creation was successful or already exists.
*/
func EnsureStorageExists() bool {
	// if storage exists but doesn't contain the 91 and 93 bytes at start and end it will overwrite.
	if !doesStorageExist() {
		fmt.Println("Storage doesn't exist! - Creating.")
		os.Mkdir("data", 0755) // create data folder, with 0755 perms (Owner: rwx, Everyone Else: r-x)

		fmt.Println(filePath)

		// create a tasks.json file in data directory
		// file contents stored in an empty array in bytes
		err = os.WriteFile(filePath, []byte("[]"), 0755)

		fmt.Println("Storage Made. - Path - ", filePath)

		if err != nil {
			fmt.Println(err)
			return false
		}

		return true
	} else {
		fmt.Println(filePath)
		return true
	}
}

func SaveNewTask(t Task) {
	if !EnsureStorageExists() {
		return
	}
	var jsonBytes []byte = encodeJSONTask(t)
	if data, err := os.ReadFile(filePath); err == nil {
		// Remove Last Element of Slice Source: https://stackoverflow.com/questions/26172196/how-to-remove-the-last-element-from-a-slice
		data = data[:len(data)-1] // Remove ]
		if data[len(data)-1] == '}' {
			data = append(data, []byte(", ")...) // ,
		}
		data = append(data, jsonBytes...) // add JSON
		data = append(data, ']')          // ] end the array.
		os.WriteFile(filePath, data, os.ModeAppend.Type().Perm())
		return
	} else {
		return
	}
}
