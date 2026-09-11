package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	// [0] is the path to the executable.
	args := os.Args[1:]

	// Display all Command-Line Args
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}
