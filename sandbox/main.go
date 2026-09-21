package main

import (
 "fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	// Access struct fields with dot
	v := Vertex{4, 2}

	// Struct pointer
	p := &v 

	// We can use without explicit dereference: p.X instead of (*p).X
	fmt.Println("The x coordinate: ", p.X)
	fmt.Println("The y coordinate: ", p.Y)

	/*
	a, b := 42, 0

	var p1 *int = &a
	fmt.Printf("Var a (value: %v) at address: %p\n", a, p1)

	b = *p1
	fmt.Printf("Var b after copying from dereferenced pointer to a: %v\n", b)
	*/
}
