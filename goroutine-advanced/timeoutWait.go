package main

import (
	"fmt"
	"sync"
	"time"
)

func timetout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		defer close(temp)
		w.Wait()
	}()

	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	var w sync.WaitGroup
	w.Add(1)

	t := 2 * time.Second
	fmt.Printf("Timeout period is %s\n", t)

	if timetout(&w, t) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}

	w.Done()
	if timetout(&w, t) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}
}
