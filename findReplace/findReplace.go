package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	// create an array of strings of size 3
	var s [3]string
	s[0] = "1 b 3"
	s[1] = "11 a B 14 1 1"
	s[2] = "b 2 -3 B -5"

	// create a regexp.Regexp if the regex is valid
	parse, err := regexp.Compile("[bB]")

	// if it is not valid, output an error and exit
	if err != nil {
		fmt.Printf("Error compiling RE: %s\n", err)
		os.Exit(-1)
	}

	// walk through the s array
	for i := 0; i < len(s); i++ {
		// replace all occurrences of b and B with C
		temp := parse.ReplaceAllString(s[i], "C")
		// output the changed string
		fmt.Println(temp)
	}

	// we can also range through the s array
	for _, val := range s {
		// replace all occurrences of b and B with C
		temp := parse.ReplaceAllString(val, "C")
		// output the changed string
		fmt.Println(temp)
	}
}
