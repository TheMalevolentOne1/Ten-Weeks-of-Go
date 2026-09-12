package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

const file = "data/tasks.json"

var dir, err = os.Getwd()
var file_path string = filepath.Join(dir, file)

func Read() ([]byte, error) {
	// Verify the path to File/Directory does not have Stats
	// and does Error i.e. Does **Not** Exist.
	if _, err := os.Stat(file_path); err != nil {
		os.Mkdir("data", 0755) // create data folder, with 0755 perms (Owner: rwx, Everyone Else: r-x)

		fmt.Println(file_path)

		// create a tasks.json file in data directory
		// file contents stored in an empty byte array
		os.WriteFile(file_path, []byte("{}"), 0755)
	}

	dataStored, err := os.ReadFile(file_path)

	if err != nil {
		return []byte{}, err
	} else {
		return []byte(dataStored), err
	}
}
