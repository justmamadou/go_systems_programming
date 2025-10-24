package main

import (
	"fmt"
	"strconv"
)

func main() {
	myArray := [4]int{2, -4, 6, 3}
	aMap := make(map[string]int)

	length := len(myArray)
	for i := 0; i < length; i++ {
		aMap[strconv.Itoa(i)] = myArray[i]
	}

	for key, value := range aMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
