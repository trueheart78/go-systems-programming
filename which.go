package main

import (
	"flag"
	"fmt"
	"os"
	"strings" // used for splitting the PATH environment var
)

func main() {
	// define the flags, and default them to false
	minusA := flag.Bool("a", false, "a")
	minusS := flag.Bool("s", false, "s")

	flag.Parse()
	flags := flag.Args()
	if len(flags) == 0 {
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}
	file := flags[0]
	found := false

	// load the PATH env var and split it
	path := os.Getenv("PATH")
	pathSlice := strings.Split(path, ":")
	for _, directory := range pathSlice {
		fullPath := directory + "/" + file
		// find out if it exists using os.Stat
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			// check the mode to see if it is a regular file, as we're not looking for symlinks or dirs
			mode := fileInfo.Mode()
			if mode.IsRegular() {
				if mode&0111 != 0 {
					found = true
					if *minusS == true { // notice the check against the pointer
						os.Exit(0)
					}
					if *minusA == true { // notice the check against the pointer
						fmt.Println(fullPath)
					} else {
						fmt.Println(fullPath)
						os.Exit(0)
					}
				}
			}
		}
	}
	// if it isn't found, exit with a status code of 1
	if found == false {
		os.Exit(1)
	}
}
