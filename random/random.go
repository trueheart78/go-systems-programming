package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// generate a random integer between the min and max
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	MIN := 0
	MAX := 0
	TOTAL := 0
	// we need more than 3 arguments, the 0 being the binary
	if len(os.Args) > 3 {
		MIN, _ = strconv.Atoi(os.Args[1])
		MAX, _ = strconv.Atoi(os.Args[2])
		TOTAL, _ = strconv.Atoi(os.Args[3])
	} else {
		// prints "Usage: [binary] MIN MAX TOTAL
		fmt.Println("Usage:", os.Args[0], "MIN MAX TOTAL")
		os.Exit(-1)
	}

	// use the epoch as the seed
	rand.Seed(time.Now().Unix())

	// createe TOTAL random numbers and print them out
	for i := 0; i < TOTAL; i++ {
		myrand := random(MIN, MAX)
		fmt.Print(myrand, " ")
	}
	fmt.Println()
}
