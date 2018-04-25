package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args

	pwd, err := os.Getwd()
	if err == nil {
		fmt.Println(pwd)
	} else {
		fmt.Println("Error:", err)
	}

	// at this point, return unless we have an arg that is '-P'
	if len(arguments) == 1 {
		return
	}
	if arguments[1] != "-P" {
		return
	}

	// if -P was the first argument, and the wd is a symlink, output its linked path
	fileinfo, err := os.Lstat(pwd)
	if fileinfo.Mode()&os.ModeSymlink != 0 {
		realpath, err := filepath.EvalSymlinks(pwd)
		if err == nil {
			fmt.Println(realpath)
		}
	}
}
