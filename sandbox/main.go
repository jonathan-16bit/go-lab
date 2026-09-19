package main

import (
	"fmt"
)

func main() {
	fmt.Println("Multiplication of 2 and 42 produces: ", mult(2, 42));

	str1, str2 := swap("Sweet", "Sour")
	fmt.Println(str1, str2)

	fmt.Println("Predecessor and successor of 43: ")
	fmt.Println(inc_dec(43))
}
