package main

import (
 "fmt"
)

func factorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}

	return fact
}

func main() {
	num := 8
	switch true {
	case num < 0:
		fmt.Println("Factorial undefined for negative numbers");
	case true:
		fmt.Printf("Factorial of %v: %v\n", num, factorial(num))
	default:
		fmt.Println("This should NOT print")
	}
}
