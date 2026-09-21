package main

import (
 "fmt"
)

func main() {
	// Creating slices with make()
	// len=1, cap=10
	numSlice := make([]int, 0, 10)
	fmt.Println("Before append:", numSlice)

	for i := 0; i <= 9; i++ {
		numSlice = append(numSlice, i)
	}

	fmt.Println("After append:", numSlice)
}
