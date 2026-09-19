package main

import (
	"fmt"
)

// Return types AFTER variable name
func mult(a int, b int) int {
	return a * b
}

// Multiple return values
func swap(a, b string) (string, string) {
	return b, a
}

// Named return values
func inc_dec(num int) (prev, next int) {
	prev = num - 1
	next = num + 1
	// Naked return: returns named values
	return
}

func main() {
	fmt.Println("Multiplication of 2 and 42 produces: ", mult(2, 42));

	str1, str2 := swap("Sweet", "Sour")
	fmt.Println(str1, str2)

	fmt.Println("Predecessor and successor of 43: ")
	fmt.Println(inc_dec(43))
}
