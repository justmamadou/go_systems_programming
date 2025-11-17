package main

import "fmt"

func main() {
	numbers := make(chan int, 5)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Printf("Not more space for: %d\n", i)
		}
	}

	for i := 0; i < counter*2; i++ {
		select {
		case x := <-numbers:
			fmt.Printf("%d\n", x)
		default:
			fmt.Println("Nothing more to print!")
			break
		}
	}
}
