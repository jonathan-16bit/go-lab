package main

import (
 "fmt"
)

func makeAdderN(n int) func(int) int {
	return func(k int) int {
		return n + k
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		next := a
		// RHS is evaluated using the old a and b
		a, b = b, a + b
		return next
	}
}

func demo_closures() {
	// Closures
	add10 := makeAdderN(10)
	for i := 0; i <= 10; i++ {
		fmt.Printf("Addition of 10 with %v = %v\n", i, add10(i))
	}

	// Fibonacci
	f := fibonacci()
	for i := 0; i <= 10; i++ {
		fmt.Println(f())
	}
}
