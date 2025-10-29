package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 0 {
		fmt.Println("Please provide an argument!")
	}
	file := arguments[1]
	err := os.Remove(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
