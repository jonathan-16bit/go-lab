package main

import (
	"fmt"
)

func main() {
	var MaxInt uint64 = 1 << 64 - 1
	var zComplex complex128 = 1i

	fmt.Println("Max int: ", MaxInt)
	fmt.Println("Square root of -1: ", zComplex)
}
