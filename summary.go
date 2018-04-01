package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// define an array of size 3
	var s [3]string

	// populate the array with space-separated data (aka columns)
	s[0] = "1 b 3"
	s[1] = "11 a 1 14 1 1"
	s[2] = "-1 2 -3 -4 -5"

	// grab the arguments passed int
	arguments := os.Args

	// convert the argument passed in to an int
	column, err := strconv.Atoi(arguments[1])

	// exit when the argument can't be converted to an int
	if err != nil {
		fmt.Println("Error reading argument")
		os.Exit(-1)
	}

	// if the column requested is 0, also exit
	if column == 0 {
		fmt.Println("Invalid column")
		os.Exit(1)
	}

	// define a sum var of 0
	sum := 0

	// cycle through the array
	for i := 0; i < len(s); i++ {
		// split the line into "columns"
		data := strings.Fields(s[i])
		// if the data contains the column
		if len(data) >= column {
			// convert to an int
			temp, err := strconv.Atoi(data[column-1])
			if err == nil {
				// add it if valid
				sum = sum + temp
			} else {
				// if the conversion was bad, tell the user
				fmt.Printf("Invalid argument: %s\n", data[column-1])
			}
		} else {
			// if the column isn't in the record, tell the user
			fmt.Println("Invalid column!")
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}
