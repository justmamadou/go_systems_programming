package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	filename := os.Args[1]
	byteSlice := []byte("Mamadou BA")
	ioutil.WriteFile(filename, byteSlice, 0644)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	defer f.Close()

	anotherByteSlice := make([]byte, 100)

	n, err := f.Read(anotherByteSlice)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Read %d bytes: %s", n, anotherByteSlice)

}
