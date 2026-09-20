package main

import (
	"fmt"
)

func main() {
	var MaxInt uint64 = 1 << 64 - 1
	var zComplex complex128 = 1i

	// Zero values for non-initialized vars
	var i int
	var f float64
	var b bool
	var s string
	// %v: default format, %q: in quotes
	fmt.Printf("%v, %v, %v, %q\n", i, f, b, s)

	fmt.Println("Max int: ", MaxInt)
	fmt.Println("Square root of -1: ", zComplex)
}
