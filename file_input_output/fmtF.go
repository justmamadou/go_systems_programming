package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	filename := os.Args[1]
	destination, err := os.Create("test")
	if err != nil {
		fmt.Printf("Eroor while createing destination file %s", err.Error())
		os.Exit(1)
	}
	defer destination.Close()

	fmt.Fprintf(destination, "[%s]:", filename)
	fmt.Fprintf(destination, "using Fprintf")
}
