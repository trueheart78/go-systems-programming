package main

import (
	"flag"
	"fmt"
)

func main() {
	// we're defining flags here, which default to false
	minusO := flag.Bool("o", false, "o")
	minusC := flag.Bool("c", false, "c")
	// except for this one, that requires an int
	minusK := flag.Int("k", 0, "an int")

	// parse the flags, as this will raise an error if an unsupported flag is passed in
	flag.Parse()

	// output the values of the flags - reminder they are addresses aka pointers
	fmt.Println("-o:", *minusO)
	fmt.Println("-c:", *minusC)
	fmt.Println("-k:", *minusK)

	// access the unused arguments passed in
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}
