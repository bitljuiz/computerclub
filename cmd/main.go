package main

import (
	"computerclub/internal/executor"
	"computerclub/internal/util"
	"fmt"
	"os"
)

// Accepts a file in 'filename' format as input
// Then it runs the executor.Execute with the given file and prints:
// 1) Line, where the error was detected
// 2) Error message and description
func main() {
	if len(os.Args)-1 != 1 {
		fmt.Printf("Invalid argument count.%sExpected: %d.%sProvided: %d%s", util.LineSeparator(),
			1, util.LineSeparator(), len(os.Args)-1, util.LineSeparator())
		return
	}
	filepath := os.Args[1]
	msg, err := executor.Execute(filepath)
	if err != nil {
		if msg != "" {
			fmt.Println(msg)
		}
		fmt.Println(err)
	}
}
