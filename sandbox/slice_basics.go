package main

import (
 "fmt"
)

func slice_basics() {
	// [N]T : N values of type T
	// Array length is part of type, so resizing is not possible
	var digits [10]int 

	for i := 0; i <= 9; i++ {
		digits[i] = i
	}

	fmt.Println(digits)

	// Slices are in the format [low: high] (half-open range)
	// Default: 0 for low, length for high
	var lower []int = digits[0:5]
	var upper []int = digits[5:10]

	fmt.Println(lower, upper)

	// Slices are references to arrays
	nums := [4]int {1, 3, 2, 4}
	a := nums[0:3]
	b := nums[1:4]
	fmt.Println(a, b)

	a[1] = 2
	fmt.Println(a, b)

	b[1] = 3
	fmt.Println(a, b)

	// Slice literals
	anotherSlice := []int {2, 3, 5, 7, 11}
	fmt.Println(anotherSlice)

	primes := [20]int{
    2, 3, 5, 7, 11,
    13, 17, 19, 23, 29,
    31, 37, 41, 43, 47,
    53, 59, 61, 67, 71,
	}

	// len() : number of elts in slice
	// cap() : number of elts in underlying array
	primeSlice := primes[:10]
	fmt.Println(primeSlice, len(primeSlice), cap(primeSlice))

	// Zero value of slice: nil
	var nilSlice []int
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nil slice")
	}
}

func appendSlice() {
	// Creating slices with make()
	// len=1, cap=10
	numSlice := make([]int, 0, 10)
	fmt.Println("Before append:", numSlice)

	for i := 0; i <= 9; i++ {
		numSlice = append(numSlice, i)
	}

	fmt.Println("After append:", numSlice)
}

func rangeForm() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	// range form: iterate over indices and values
	for i, v := range pow {
		fmt.Printf("2**%v = %v\n", i, v)
	}
}
