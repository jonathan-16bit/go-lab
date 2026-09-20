package main

import (
	"fmt"
)

func main() {
	// Type inference
	var w int = 42
	v := w
	fmt.Printf("v (%v) is of type: %T\n", v, v)

	/*
	// Type conversions
	var fl float64 = -1.2
	var x int = int(fl)
	var u uint64 = uint64(x)
	fmt.Println(fl, x, u)

	// Zero values for non-initialized vars
	var i int
	var f float64
	var b bool
	var s string
	// %v: default format, %q: in quotes
	fmt.Printf("%v, %v, %v, %q\n", i, f, b, s)

	// Types
	var MaxInt uint64 = 1 << 64 - 1
	var zComplex complex128 = 1i
	fmt.Println("Max int: ", MaxInt)
	fmt.Println("Square root of -1: ", zComplex)
	*/
}
