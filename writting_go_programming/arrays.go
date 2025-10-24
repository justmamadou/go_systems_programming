package main

import "fmt"

func main() {
	myArray := [4]int{1, 2, 4, -4}
	twoD := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	threeD := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	for _, v := range myArray {
		fmt.Printf("%d ", v)
	}

	fmt.Println()
	for _, v := range twoD {
		fmt.Printf("%d ", v)
	}

	fmt.Println()
	for _, v := range threeD {
		fmt.Printf("%d ", v)
	}

}
