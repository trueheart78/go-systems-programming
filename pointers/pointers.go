package main

import (
	"fmt"
)

// requires no return value since we're changing the pointer reference
func withPointer(x *int) {
	*x = *x * *x
}

type complex struct {
	x, y int
}

// requires a return value
func newComplex(x, y int) *complex {
	return &complex{x, y}
}

func main() {
	x := -2
	// pass by pointer reference requires a leading '&'
	withPointer(&x)
	// will print the changed value: 4
	fmt.Println(x)

	w := newComplex(4, -5)
	// will print the actual data: {4 -5}
	fmt.Println(*w)
	// will print the reference to the data: &{4 -5}
	fmt.Println(w)
}
