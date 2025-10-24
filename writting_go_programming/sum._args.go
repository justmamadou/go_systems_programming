package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0
	argments := os.Args
	for i := 1; i < len(argments); i++ {
		valeur, _ := strconv.Atoi(argments[i])
		sum = sum + valeur
	}
	fmt.Printf("Sum: %d\n", sum)
}
