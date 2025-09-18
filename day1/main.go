package main

import (
	"fmt"
	"myproject/src" // import using module name + folder
)

// arithmetic operations
func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

// main function
func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Add:", add(5, 3))
	src.PrintUser()
}
