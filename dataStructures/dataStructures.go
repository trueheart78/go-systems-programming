package main

import (
	"fmt"
	"reflect"
)

func main() {

	// message has 3 fields: X, Y, and Label
	type message struct {
		X     int
		Y     int
		Label string
	}

	// initialize with passed values
	p1 := message{23, 12, "A Message"}

	// initialize with zero values
	p2 := message{}

	// assign a new value
	p2.Label = "Message 2"

	s1 := reflect.ValueOf(&p1).Elem()
	s2 := reflect.ValueOf(&p2).Elem()

	// prints: S2= {0 0 Message 2}
	fmt.Println("S2=", s2)

	typeOfT := s1.Type()
	// prints: P1= {23 12 A Message}
	fmt.Println("P1=", p1)
	// prints: P2= {0 0 Message 2}
	fmt.Println("P2=", p2)

	for i := 0; i < s1.NumField(); i++ {
		f := s1.Field(i)

		// print the value using _any_ type
		fmt.Printf("%d: %s ", i, typeOfT.Field(i).Name)
		fmt.Printf("%s = %v\n", f.Type(), f.Interface())
	}
}
