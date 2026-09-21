package main

import (
 "fmt"
)

func main() {
	a, b := 42, 0

	var p1 *int = &a
	fmt.Printf("Var a (value: %v) at address: %p\n", a, p1)

	b = *p1
	fmt.Printf("Var b after copying from dereferenced pointer to a: %v\n", b)
}
