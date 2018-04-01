package main

import (
	"fmt"
	"strings"
)

func main() {
	// define an array of size 3 that has strings
	var s [3]string
	s[0] = "1 b 3 1 a a b"
	s[1] = "11 a 1 1 1 1 a a"
	s[2] = "-1 b 1 -4 a 1"

	// create a map (aka hash), where keys are strings and values are ints
	// make(map[key type]value type)
	counts := make(map[string]int)

	// populate the map with data from the array
	for i := 0; i < len(s); i++ {
		// split the string on whitespace for columns
		data := strings.Fields(s[i])
		// range over the data, ignoring the key but grabbing the value
		for _, word := range data {
			// attempt to assign the the counts map
			_, ok := counts[word]
			if ok {
				// if good, increments
				counts[word] = counts[word] + 1
			} else {
				// otherwise, initialize
				counts[word] = 1
			}
		}
	}

	// range over it and output the counts
	for key, _ := range counts {
		fmt.Printf("%s -> %d \n", key, counts[key])
	}
}
