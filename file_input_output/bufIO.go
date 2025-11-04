package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error while reading %s %s", &filename, err.Error())
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if scanner.Err() != nil {
			fmt.Printf("error reading file %s", err)
			os.Exit(1)
		}
		fmt.Println(line)
	}

}
