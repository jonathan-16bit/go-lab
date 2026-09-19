package main

import (
	"fmt"
)

// Return types AFTER variable name
func mult(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println("Multiplication of 2 and 42 produces: ", mult(2, 42));
}
