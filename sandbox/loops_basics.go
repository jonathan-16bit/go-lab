package main

import (
	// "fmt"
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

	for {
		if i > n {
			break
		}

		fact *= i
		i += 1
	}

	return fact
}
