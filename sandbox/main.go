package main

import (
 "fmt"
)

func main() {
	fmt.Println("Go experiment")

	for i := 0; i <= 9; i++ {
		defer fmt.Printf("%v\n", i)
	}

	fmt.Println("Count down: ")
}
