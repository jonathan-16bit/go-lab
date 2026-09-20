package main

import (
 "fmt"
 "math"
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
	if num == 0 {
		fmt.Println("Factorial of 0: 1")
		return
	}

	sign := int(math.Abs(float64(num))) / num

	switch sign {
	case -1:
		fmt.Println("Factorial undefined for negative numbers");
	case 1:
		fmt.Printf("Factorial of %v: %v\n", num, factorial(num))
	default:
		fmt.Println("This should NOT print")
	}
}
