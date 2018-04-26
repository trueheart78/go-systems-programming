package main

import (
	"fmt"
	"reflect"
)

func main() {
	// create two custom types: t1 and t2
	type t1 int
	type t2 int

	// create 3 different int variables to reflect on
	x1 := t1(1)
	x2 := t2(1)
	x3 := 1

	// find the concrete values of the 3 different vars
	st1 := reflect.ValueOf(&x1).Elem()
	st2 := reflect.ValueOf(&x2).Elem()
	st3 := reflect.ValueOf(&x3).Elem()

	// print the concrete values
	fmt.Printf("ST1: %v\n", st1)
	fmt.Printf("ST2: %v\n", st2)
	fmt.Printf("ST3: %v\n", st3)

	// detect the types of each of the 3 different vars
	fmt.Printf("X1 Type: %s\n", st1.Type())
	fmt.Printf("X2 Type: %s\n", st2.Type())
	fmt.Printf("X3 Type: %s\n", st3.Type())

	// create another, more complex type
	type aStructure struct {
		X    uint
		Y    float64
		Text string
	}

	// create a var to reflect on
	x4 := aStructure{123, 3.14, "A Structure"}

	// find the concrete values of the new var
	st4 := reflect.ValueOf(&x4).Elem()
	fmt.Printf("ST4: %v\n", st4)

	// detect the type of the new var
	typeOfX4 := st4.Type()

	// print the type of the new var
	fmt.Printf("X4 Type: %s\n", typeOfX4)
	fmt.Printf("The fields of %s are:\n", typeOfX4)

	// cycle through the number of fields in the type and output their type and value
	for i := 0; i < st4.NumField(); i++ {
		fmt.Printf("%d: Field name: %s ", i, typeOfX4.Field(i).Name)
		fmt.Printf("Type: %s ", st4.Field(i).Type())
		fmt.Printf("and Value: %v\n", st4.Field(i).Interface())
	}
}
