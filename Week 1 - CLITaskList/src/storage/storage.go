package storage

// Storage Module.
// Brief: Writes and Reads the Tasks Struct Marshaling and Unmarshaling the JSON Encoding from the bytes which are Write/Read to/from the file.

// IMPORTANT NOTE TO SELF
/*
Read and Write handle in BYTES for I/O Separation
*/

import (
	"encoding/json"  // For JSON handling
	"fmt"            // Format for console display
	"os"             // OS for arguments
	"path/filepath"  // For Filepath handling
	"tasklist/tasks" // For the Tasks Struct
)

const file = "data/tasks.json"

var dir, err = os.Getwd() // GetWorkingDirectory

var file_path string = filepath.Join(dir, file) // json data file path

var tmp_tasks []tasks.Task // Empty Tasks Array. Prevents Editing/Corruption of Tasks Module List.
var tmp_task tasks.Task

/*
Brief: Encodes Task Struct from Struct into JSON
Parameters:
b - bytes (the bytes of the Tasks slice)
*/
func encodeJSONTask(t tasks.Task) []byte {
	data, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	} else {
		return data
	}
}

/*
Brief: Decodes bytes of JSON into the empty Tasks Struct
Parameters:
b - bytes (the bytes of JSON data)
*/
func decodeJSONTasks(b []byte) {
	err = json.Unmarshal(b, &tmp_task)
	if err != nil {
		fmt.Println(err)
	} else {
		return
	}
}

func doesStorageExist() bool {
	if _, err := os.Stat(file); err == nil {
		return true
	} else {
		return false
	}
}

func EnsureStorageExists() bool {
	if !doesStorageExist() {
		fmt.Println("Storage doesn't exist! - Creating.")
		os.Mkdir("data", 0755) // create data folder, with 0755 perms (Owner: rwx, Everyone Else: r-x)

		fmt.Println(file_path)

		// create a tasks.json file in data directory
		// file contents stored in an empty array in bytes
		err = os.WriteFile(file_path, []byte("[]"), 0755)

		fmt.Println("Storage Made. - Path - ", file_path)

		if err != nil {
			fmt.Println(err)
			return false
		}

		return true
	} else {
		fmt.Println(file_path)
		return true
	}
}

// SHOULD BE COMPLETELY REDONE TO FIT INTO PLAN
// func Read() ([]byte, error) {}
// func Write([]byte) (error) {}

func SaveTask(t tasks.Task) {
	jsonT := encodeJSONTask(t)
	fmt.Println(jsonT)
}
