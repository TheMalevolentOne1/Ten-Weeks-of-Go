/*
Brief: Storage Package
Handles json file storage returning byte array and error.
*/
package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const file = "data/tasks.json"

var dir, err = os.Getwd()

var file_path string = filepath.Join(dir, file)

func FileExists(f string) bool {
	if _, err := os.Stat(f); err == nil {
		return true
	} else {
		return false
	}
}

func storageExists() {
	if !FileExists(file_path) {
		os.Mkdir("data", 0755) // create data folder, with 0755 perms (Owner: rwx, Everyone Else: r-x)

		fmt.Println(file_path)

		// create a tasks.json file in data directory
		// file contents stored in an empty byte array
		err = os.WriteFile(file_path, []byte("[]"), 0755)

		if err != nil {
			fmt.Println(err)
		}

		return
	} else {
		return
	}
}

func Read() ([]byte, error) {
	storageExists()
	dataStored, err := os.ReadFile(file_path)

	json.Unmarshal(dataStored)

	if err != nil {
		return []byte{}, err
	} else {
		return []byte(dataStored), err
	}
}

func Write(task []byte) bool {

}
