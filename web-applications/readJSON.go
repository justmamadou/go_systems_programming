package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}
type Telephone struct {
	Mobile bool
	Number string
}

func LoadFromJSON(filename string, key interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	decodeJSON := json.NewDecoder(f)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		os.Exit(100)
	}
	filename := arguments[1]

	var myRecord Record
	err := LoadFromJSON(filename, &myRecord)
	if err == nil {
		fmt.Println(myRecord)
	} else {
		fmt.Println(err)
	}

}
