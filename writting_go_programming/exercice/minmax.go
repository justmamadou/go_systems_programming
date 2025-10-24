package main

import (
	"fmt"
)

func findMinMax(numbers []int) {
	valeur := 10
	for valeur != 0 {
		fmt.Println("Enter a number: ")
		fmt.Scan(&valeur)

		numbers = append(numbers, valeur)
	}
	index := len(numbers) - 1
	numbers = append(numbers[:index], numbers[index+1:]...)

	max := 0
	min := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
		if numbers[i] < min && min != 0 {
			min = numbers[i]
		}
	}
	fmt.Printf("Max=%d, Min=%d\n", max, min)
}

/*
	func findFromCMD() {
		var numbers []int
		max := 0
		min := 0
		arguments := os.Args
		if len(arguments) < 2 {
			fmt.Println("Usage: go run minmax.go 1 2 3")
			return
		}
		for i := 1; i < len(arguments); i++ {
			valeur, _ := strconv.Atoi(arguments[i])
			numbers = append(numbers, valeur)
			if numbers[i] > max {
				max = numbers[i]
			}
			if numbers[i] < min {
				min = numbers[i]
			}
		}
		fmt.Printf("Max=%d, Min=%d\n", max, min)

}
*/
func main() {
	//var numbers []int
	//findMinMax(numbers)
	findFromCMD()
}
