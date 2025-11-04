package main

import (
	"fmt"
	"io"
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
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	buf := make([]byte, 8)
	if _, err := io.ReadFull(f, buf); err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}
	io.WriteString(os.Stdout, string(buf))
	fmt.Println()
}
