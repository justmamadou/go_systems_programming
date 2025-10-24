package main

import "fmt"

type coordinates interface {
	xaxis() int
	yaxis() int
}

type point2D struct {
	X int
	Y int
}

func (p point2D) xaxis() int {
	return p.X
}

func (p point2D) yaxis() int {
	return p.Y
}

func findCoordinates(c coordinates) {
	fmt.Printf("X-Axis=%d and Y-Axis=%d\n", c.xaxis(), c.yaxis())
}

type xpoint int
type ypoint int

func (x xpoint) xaxis() int {
	return int(x)
}

func (x xpoint) yaxis() int {
	return 0
}

func (y ypoint) xaxis() int {
	return 0
}

func (y ypoint) yaxis() int {
	return int(y)
}

func main() {
	p := point2D{X: -1, Y: 12}
	fmt.Println(p)
	findCoordinates(p)
	x := xpoint(10)
	findCoordinates(x)
	y := ypoint(20)
	findCoordinates(y)
}
