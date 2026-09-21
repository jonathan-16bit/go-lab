package main

import (
 "fmt"
)

func main() {
	// [N]T : N values of type T
	// Array length is part of type, so resizing is not possible
	var digits [10]int 

	for i := 0; i <= 9; i++ {
		digits[i] = i
	}

	fmt.Println(digits)
}
