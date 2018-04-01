package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value int64 = 5

	// define p1 as an int64 pointer value
	var p1 = &value
	// use the unsafe lib to reference p1 as a different type than
	// it was defined as
	var p2 = (*int32)(unsafe.Pointer(p1))

	// they are both happy with 5
	fmt.Println("*p1: ", *p1)
	fmt.Println("*p2: ", *p2)
	// change the pointer to be an int64
	*p1 = 312121321321213212

	// it is stored fine when referenced as an int64
	fmt.Println(value)
	// this reference in an int32 type outputs 606940444 which is totes incorrect due to the data type
	// this is the unsafe pointer operation and you can see why since the variable is an unexpected value
	fmt.Println("*p2: ", *p2)

	//change the pointer to be referenceable by an int32
	*p1 = 31212132

	// now they both output the same value
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
}
