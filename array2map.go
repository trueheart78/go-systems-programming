package main

import (
	"fmt"
	"strconv"
)

func main() {
	anArray := [4]int{1, -2, 14, 0}
	aMap := make(map[string]int)

	length := len(anArray)

	// visit all elements of the array and add them to the map
	for i := 0; i < length; i++ {
		fmt.Printf("%s ", strconv.Itoa(i))
		// key is the int converted to a string, value is the int itself
		aMap[strconv.Itoa(i)] = anArray[i]
	}

	// print the mapped values using a range (think foreach)
	for key, value := range aMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
