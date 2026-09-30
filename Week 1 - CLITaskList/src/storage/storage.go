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

var dir, err = os.Getwd()

var file_path string = filepath.Join(dir, file)

var tmp_tasks_copy []tasks.Task // Empty Tasks Array. Prevents Editing/Corruption of Tasks Module List.

func encodeJSON(b []byte) []byte {
	data, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	} else {
		return data
	}
}

func decodeJSON(b []byte) {
	err = json.Unmarshal(b, &tmp_tasks_copy)
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

func Read() {

}
