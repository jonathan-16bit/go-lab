package main

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
