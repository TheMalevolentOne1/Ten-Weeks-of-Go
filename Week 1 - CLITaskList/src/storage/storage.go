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

var filePath string = filepath.Join(dir, file) // json data file path

var tmpTasks []tasks.Task // Empty Tasks Array. Prevents Editing/Corruption of Tasks Module List.
var tmpTask tasks.Task

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
	err = json.Unmarshal(b, &tmpTasks)
	if err != nil {
		fmt.Println(err)
	} else {

	}
}

// DOES THE FILE HAVE THE CORRECT BYTES [91, 93] representing the empty array []. Can the correct storage pattern by identified by the byte pattern? or in a ReGEx.
func doesStorageExist() bool {
	if _, err := os.Stat(file); err == nil {
		return true
	} else {
		return false
	}
}

/*
Brief: Ensure the storage exists, and create it.
Returnn Boolean whether creation was successful or already exists.
*/
func EnsureStorageExists() bool {
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

func SaveTask(t tasks.Task) {
	if !EnsureStorageExists() {
		return
	}
	var jsonBytes []byte = encodeJSONTask(t)
	if data, err := os.ReadFile(filePath); err == nil {
		// Remove Last Element of Slice Source: https://stackoverflow.com/questions/26172196/how-to-remove-the-last-element-from-a-slice
		data = data[:len(data)-1]            // Remove [
		data = append(data, jsonBytes...)    // add JSON
		data = append(data, []byte(", ")...) // ,
		data = append(data, '\n')            // newline
		data = append(data, ']')             // ] end the array.
		os.WriteFile(filePath, data, os.ModeAppend.Type().Perm())
		return
	} else {
		return
	}
}
