package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Example function returning error
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Example with strconv conversion
func stringToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("conversion error: %v", err)
	}
	return num, nil
}

func main() {
	// Divide example
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// String conversion example
	number, err := stringToInt("123a")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted number:", number)
	}

	// Successful case
	success, err := divide(20, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("20 / 5 =", success)
	}
}
