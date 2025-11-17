package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var aMutex sync.Mutex
var sharedVariable string = ""

func addDot() {
	aMutex.Lock()
	sharedVariable = sharedVariable + "."
	aMutex.Unlock()
}

func read() string {
	aMutex.Lock()
	res := sharedVariable
	aMutex.Unlock()
	return res
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s numberOfGoRoutine", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	var waitGoup sync.WaitGroup
	numGR, _ := strconv.ParseInt(os.Args[1], 10, 64)
	for i := 0; i < int(numGR); i++ {
		waitGoup.Add(1)
		go func() {
			defer waitGoup.Done()
			addDot()
		}()
	}

	waitGoup.Wait()
	fmt.Printf("->%s\n", read())
	fmt.Printf("Length: %d\n", len(read()))
}
