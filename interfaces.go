package main

import (
	"fmt"
)

// declare an interface that has two methods required
type coordinates interface {
	xaxis() int
	yaxis() int
}

// declare a new type called point2D
type point2D struct {
	X int
	Y int
}

// create the xaxis() method on the point2D type
func (s point2D) xaxis() int {
	return s.X
}

// create the yaxis() method on the point2D type
func (s point2D) yaxis() int {
	return s.Y
}

// call the xaxis() and yaxis() method on the passed in coordinates interface
// notice this is not a soloCoordinate or point2D type, it is the interface
// defined further above
func findCoordinates(a coordinates) {
	fmt.Println("X:", a.xaxis(), "Y:", a.yaxis())
}

// declare a new type called coordinate
type soloCoordinate int

// create the xaxis() method on the soloCoordinate type
// and return the integer value of the related type
func (s soloCoordinate) xaxis() int {
	return int(s)
}

// create the yaxis() method on the soloCoordinate type
// and return 0
func (s soloCoordinate) yaxis() int {
	return 0
}

func main() {
	// new point2D instance
	x := point2D{X: -1, Y: 12}
	fmt.Println(x)

	// send in the point2D var - prints X: -1 Y: 12
	findCoordinates(x)

	y := soloCoordinate(10)
	// send in the point2D var - prints X: 10 Y: 0
	findCoordinates(y)
}
