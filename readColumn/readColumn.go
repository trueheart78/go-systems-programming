package main

import (
	"fmt"
	"strings"
)

func main() {
	// define an array of size 3
	var s [3]string

	// populate the array with space-separated values
	s[0] = "1 2 3"
	s[1] = "11 12 13 14 15 16"
	s[2] = "-1 2 -3 -4 -5 6"

	// define the column to read
	column := 2

	// go through each element of the array
	for i := 0; i < len(s); i++ {
		// split the string on whitespace characters
		data := strings.Fields(s[i])
		// if the data size has space for the relevant column
		if len(data) >= column {
			// print out the column
			fmt.Println((data[column-1]))
		}
	}
}
