package main

import (
 "fmt"
)

type Vertex struct {
	X, Y int
}

func main() {
	// Order of named fields is irrelevant
	v1 := Vertex{4, 2}
	v2 := Vertex{Y: 3, X: 4}

	v4 := Vertex{Y: 1}  // X: 0
	v3 := Vertex{}  // Both are 0

	v5 := &Vertex{4, 4}

	fmt.Println(v1, v2, v3, v4, v5)
}
