package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument")
	}
	file := arguments[1]

	fileInfo, err := os.Stat(file)
	if err != nil {
		fmt.Printf("Error while getting file info %s", err.Error())
		os.Exit(1)
	}
	mode := fileInfo.Mode()
	fmt.Println(file, "\nmode: ", mode)
}
