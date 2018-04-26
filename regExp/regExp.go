package main

import (
	"fmt"
	"regexp"
)

func main() {
	// matches
	match, _ := regexp.MatchString("Mihalis", "Mihalis Tsoukalos")
	fmt.Println(match)
	// does not match, case-sensitive
	match, _ = regexp.MatchString("Tsoukalos", "Mihalis tsoukalos")
	fmt.Println(match)

	// validate the regex and return a usable regex.Regexp var
	// that can perform matches
	parse, err := regexp.Compile("[Mm]ihalis")

	if err != nil {
		fmt.Printf("Error compiling RE: %s\n", err)
	} else {
		fmt.Println(parse.MatchString("Mihalis Tsoukalos"))
		fmt.Println(parse.MatchString("mihalis Tsoukalos"))
		fmt.Println(parse.MatchString("M ihalis Tsoukalos"))

		// search the provided string ("mihalis Mihalis") for matches, and replace them with "MIHALIS"
		fmt.Println(parse.ReplaceAllString("mihalis Mihalis", "MIHALIS"))
	}
}
