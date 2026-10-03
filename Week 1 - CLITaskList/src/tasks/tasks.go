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

const file = "data/tasks.json"

var dir, err = os.Getwd() // GetWorkingDirectory

var filePath string = filepath.Join(dir, file) // json data file path

func GetNextTaskID() int {
	fmt.Println(Tasks)
	return len(Tasks) + 1
}

func ShowAllTasks() {
	if lenTasks := len(Tasks); lenTasks == 0 {
		fmt.Println("No Tasks.")
	} else {
		for i := 0; i < lenTasks; i++ {
			fmt.Printf("Task Number: %d \nTask: %s \nDescription: %s \n",
				Tasks[i].Id,
				Tasks[i].Name,
				Tasks[i].Description,
			)
		}
	}
}

/*
Brief: Encodes Task Struct from Struct into JSON
Parameters:
b - bytes (the bytes of the Tasks slice)
*/
func encodeJSONTask(t Task) ([]byte, error) {
	data, err := json.Marshal(t)

	if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

/*
DecodeAllJsonTasks ensures that all JSON Tasks decoded from the JSON File into the Tasks Struct otherwise returns error.
Parameters:
b - bytes (the bytes of JSON data)
*/
func DecodeAllJSONTasks() error { // Load JSON File into Tasks Struct Memory
	if b, err := os.ReadFile(filePath); err == nil {
		err = json.Unmarshal(b, &Tasks) // Decode JSON into Tasks Struct
		if err != nil {
			return err
		}
	}

	return nil
}

// doesStorageExist ensures json file exists with [ bytes at the start and end ]
// returns boolean with result if exists, with valid empty slice.
func doesStorageExist() bool {
	if _, err := os.Stat(filePath); err == nil {
		if data, err := os.ReadFile(filePath); err == nil {
			// confirm length contains at least two and [ and ] and CORRECT BYTES [91, 93]
			if len(data) >= 2 && data[0] == '[' && data[len(data)-1] == ']' {
				return true
			}
		}
	}

	return false
}

var filePerm = os.ModeAppend.Type().Perm()

/*
EnsureStorageExists ensures the storage exists, and if it doesn't create it.
Returnn Boolean whether creation was successful or already exists.
*/
func EnsureStorageExists() bool {
	// if storage exists but doesn't contain the 91 and 93 bytes at start and end it will overwrite.
	if !doesStorageExist() {
		fmt.Println("Storage doesn't exist! - Creating.")
		os.Mkdir("data", filePerm) // create data folder, with 0655 perms (Owner: rwx, Everyone Else: r-x)

		fmt.Println(filePath)

		// create a tasks.json file in data directory
		// file contents stored in an empty array in bytes
		err = os.WriteFile(filePath, []byte("[]"), filePerm)

		fmt.Println("Storage Made. - Path - ", filePath)

		if err != nil {
			fmt.Println(err)
			return false
		} else {
			return true
		}
	} else {
		fmt.Println(filePath)
		return true
	}
}

func SaveNewTask(t Task) error {
	var taskEncoded []byte

	if tJSON, err := encodeJSONTask(t); err != nil {
		taskEncoded = tJSON
		return err
	}

	if data, err := os.ReadFile(filePath); err == nil {
		// Remove Last Element of Slice Source: https://stackoverflow.com/questions/26172196/how-to-remove-the-last-element-from-a-slice
		data = data[:len(data)-1]     // Remove ]
		if data[len(data)-1] == '}' { // Ensure previous is closing json tag
			data = append(data, []byte(", ")...) // ,
		}
		data = append(data, taskEncoded...) // add JSON
		data = append(data, ']')            // ] end the array.
		os.WriteFile(filePath, data, filePerm)
		return nil
	} else {
		return nil
	}
}

// TODO: ID numbers could overlap as based on size.
// - Some ID Correction function should be implemented to ensure the numbers are accurate;
// Lazier Method: Scratching that overwriting the IDs to be manually accurate 1,2,3,3,5 to 1,2,3,4,5.
// Problem: Loses integrity of the original ID... The Wrong ID could be deleted in sequential updates.
// No undo or error mitigation handling...
func fixIDNumbers() {

}

// FindID ensures that the Task ID is found and the first { position is returned.
func FindIDTask(id int) (int, error) {
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(Tasks); i++ {
		if Tasks[i].Id == id {
			return i, nil
		} else {
			continue
		}
	}

	return id, nil
}

// RemoveTask ensures a task is removed from the JSON file using the first curly bracket then deleting all subsequent bytes adding a comma if necessary.
func RemoveTask(tPos int) error {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return err
	}

	var jsonStartPos int
	var jsonEndPos int

	count := 0

	for i, jsonPos := range data {
		if count != tPos && jsonPos == '{' {
			count++
		}

		if count == tPos && jsonPos == '{' { // find the { when the right index is located for Task
			jsonStartPos = i
		}

		if count == tPos && jsonPos == '}' {
			jsonEndPos = i
			break
		}
	}

	fmt.Println(count)
	fmt.Println(jsonStartPos)
	fmt.Println(jsonEndPos)

	return nil
}
