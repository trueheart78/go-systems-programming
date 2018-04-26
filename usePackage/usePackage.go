package main

import (
	"fmt"

	"github.com/trueheart78/go-systems-programming/aSimplePackage"
)

func main() {
	temp := aSimplePackage.Add(5, 10)
	fmt.Println(temp)

	fmt.Println(aSimplePackage.Pi)
}
