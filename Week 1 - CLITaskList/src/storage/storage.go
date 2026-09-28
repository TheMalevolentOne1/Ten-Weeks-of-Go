package storage

// Storage Module.
// Brief: Writes and Reads the Tasks Struct Marshaling and Unmarshaling the JSON Encoding from the bytes which are Write/Read to/from the file.

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

func doesStorageExist() bool {
	if _, err := os.Stat(file); err == nil {
		return true
	} else {
		return false
	}
}

func EnsureStorageExists() bool {
	if !doesStorageExist() {
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
		return true
	}
}

func Read() ([]byte, error) {
	if !EnsureStorageExists() {
	}
	dataStored, err := os.ReadFile(file_path)

	json.Unmarshal(dataStored, &tasks.Tasks) // Decode JSON data from dataStored into the Tasks slice address.

	if err != nil {
		return []byte{}, err
	} else {
		return []byte(dataStored), err
	}
}
