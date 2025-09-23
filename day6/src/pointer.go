package src

import "fmt"

func PointersDemo() {
	// Basic Pointer
	var x int = 10
	var p *int = &x // p stores the address of x

	fmt.Println("x value:", x)
	fmt.Println("x address:", &x)
	fmt.Println("p value (address of x):", p)
	fmt.Println("value at p (*p):", *p)

	// Modifying through pointer

	*p = 20 // changes the value of x
	fmt.Println("\nAfter *p = 20")
	fmt.Println("x value:", x)

	// Function with pointer
	fmt.Println("\nPointer with function")
	y := 5
	fmt.Println("Before:", y)
	double(&y) // pass address
	fmt.Println("After:", y)

	// Nil pointer
	var z *int
	if z == nil {
		fmt.Println("\nz is nil (no address assigned yet)")
	}
}

func double(num *int) {
	*num = *num * 2
}
