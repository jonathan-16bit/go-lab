package main

import (
 "fmt"
)

func main() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	// range form: iterate over indices and values
	for i, v := range pow {
		fmt.Printf("2**%v = %v\n", i, v)
	}
}
