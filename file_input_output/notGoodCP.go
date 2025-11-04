package main

import (
	"fmt"
	"io"
	"os"
)

func Copy(destination, source string) (int64, error) {
	soureStat, err := os.Stat(source)
	if err != nil {
		return 0, fmt.Errorf("Error while checking source file info: %s", err.Error())
	}
	if !soureStat.Mode().IsRegular() {
		return 0, fmt.Errorf("Not a regular file: %s", err.Error())
	}
	f, err := os.Open(source)
	if err != nil {
		return 0, fmt.Errorf("Error while opening source file %s: %s", source, err.Error())
	}
	defer f.Close()
	d, err := os.Create(destination)
	if err != nil {
		return 0, fmt.Errorf("Error while creating destination file: %s", err.Error())
	}
	defer d.Close()
	nBytes, err := io.Copy(d, f)
	return nBytes, err
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide two command line arguments!")
		os.Exit(1)
	}

	sourceFiile := os.Args[1]
	destinationFile := os.Args[2]
	nBytes, err := Copy(destinationFile, sourceFiile)
	if err != nil {
		fmt.Printf("The copy operation failed %q\n", err)
	} else {
		fmt.Printf("Copied %d bytes!\n", nBytes)
	}

}
