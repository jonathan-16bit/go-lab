package main

import (
	"fmt"
	"reflect"
)

// Package level
var fact_2 = 2

func main() {
	// Function level
	var fact_3 = 6

	// Short variable declaration
	fact_4 := 24

	// Type is inferred from initializer
	fmt.Println("Type of fact_2: ", reflect.TypeOf(fact_2))
	fmt.Println("Type of fact_3: ", reflect.TypeOf(fact_3))
	fmt.Println("Type of fact_4: ", reflect.TypeOf(fact_4))
}
