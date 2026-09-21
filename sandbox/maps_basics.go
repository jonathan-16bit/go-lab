package main

import (
 "fmt"
)

func listPerfectSquares() {
	square := make(map[int]int)
	fmt.Println("Squares map before adding KV pairs:", square)

	for i := 1; i <= 10; i++ {
		square[i] = i * i
	}

	fmt.Println("Squares map after adding KV pairs:", square)

	invertSquare := make(map[int]int)
	for k, v := range square {
		invertSquare[v] = k
	}

	// Checking if values exist
	fmt.Println("Perfect squares till 100: ")
	for i := 0; i <= 100; i++ {
		_, ok := invertSquare[i]

		if ok == true {
			fmt.Println(i)
		}
	}
}
