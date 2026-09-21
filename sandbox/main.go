package main

import (
 "fmt"
)

func makeAdderN(n int) func(int) int {
	return func(k int) int {
		return n + k
	}
}

func main() {
	// Closures
	add10 := makeAdderN(10)
	for i := 0; i <= 10; i++ {
		fmt.Printf("Addition of 10 with %v = %v\n", i, add10(i))
	}
}
