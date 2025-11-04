package main

import (
	"fmt"
	"io"
	"os"
)

func countChars(r io.Reader) int {
	buf := make([]byte, 16)
	total := 0

	for {
		n, err := r.Read(buf)
		if err != nil && err!= io.EOF {
			return 0
		}
		if err == io.EOF {
			break
		}
		total = total + n
	}
	return total
}

func writeNumberChars(w io.Writer, x int){
	fmt.Fprintf(w, "%d\n",x)
}

func main(){

	if len(os.Args) != 2 {
		fmt.Println("Please provide the filename")
		os.Exit(1)
	}

	fiilename := os.Args[1]
	_,err := os.Stat(fiilename)
	if err != nil {
		fmt.Printf("Error with the file provided %s", err.Error())
		os.Exit(1)
	}
	f, err := os.Open(fiilename)
	if err != nil{
		fmt.Printf("Error while opening file: %s", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	chars := countChars(f)
	fiilename =  fiilename+".count"
	f, err = os.Create(fiilename)
	if err != nil {
		fmt.Printf("Error creating file %s: %s",fiilename, err.Error())
	}
	defer f.Close()

	writeNumberChars(f, chars)
}