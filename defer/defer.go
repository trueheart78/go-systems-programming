package main

import (
	"fmt"
)

// LIFO - will print 2 1 0
func a1() {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i, " ")
	}
}

// LIFO - the function body is run after the for loop completes
//        and results in the printing of 3 3 3
func a2() {
	for i := 0; i < 3; i++ {
		defer func() { fmt.Print(i, " ") }()
	}
}

// LIFO - since we're passing in the value, it does not use the
//        value of i after the loop completes. Will print 2 1 0
func a3() {
	for i := 0; i < 3; i++ {
		defer func(n int) { fmt.Print(n, " ") }(i)
	}
}

func main() {
	a1()
	fmt.Println()
	a2()
	fmt.Println()
	a3()
	fmt.Println()
}
