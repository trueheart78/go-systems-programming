package main

import (
	"fmt"
)

// no named return values
func unnamedMinMax(x, y int) (int, int) {
	if x > y {
		min := y
		max := x
		return min, max
	} else {
		min := x
		max := y
		return min, max
	}
}

// named return values
func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return min, max
}

// named return values with naked return values
func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return
}

// sort without using an extra variable
func sort(x, y int) (int, int) {
	if x > y {
		return x, y
	} else {
		return y, x
	}
}

func main() {
	y := 4
	// anonymous function usage
	square := func(s int) int {
		return s * s
	}
	fmt.Println("The square of", y, "is", square(y))

	square = func(s int) int {
		return s + s
	}
	fmt.Println("The square of", y, "is", square(y))

	// named function usage
	fmt.Println(minMax(15, 6))
	fmt.Println(namedMinMax(15, 6))
	min, max := namedMinMax(12, -1)
	fmt.Println(min, max)
}
