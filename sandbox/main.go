package main

import (
	"fmt"
)

func factorial_for(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}

	return fact
}

// Also for-loop, but while-style
func factorial_while(n int) int {
	fact, i := 1, 1

	for i <= n {
		fact *= i
		i += 1
	}

	return fact
}

func main() {
	n := 20
	fmt.Printf("Factorial of %v: %v\n", n, factorial_for(n))

	m := 20
	fmt.Printf("Factorial of %v: %v\n", m, factorial_while(m))
}
