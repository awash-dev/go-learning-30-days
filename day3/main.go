package main

import (
	"fmt"
)

func main() {
	// --- IF STATEMENTS ---
	age := 20
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 60 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior.")
	}

	// --- FOR LOOP ---
	fmt.Println("\nNumbers 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Range-based loop
	numbers := []int{10, 20, 30, 40}
	fmt.Println("\nUsing range:")
	for index, value := range numbers {
		fmt.Printf("Index %d => Value %d\n", index, value)
	}

	// Infinite loop with break
	fmt.Println("\nInfinite loop (break at 3):\n")
	counter := 1
	for {
		if counter > 3 {
			break
		}
		fmt.Println("Counter:", counter)
		counter++
	}

	// --- SWITCH STATEMENT ---
	day := 3
	fmt.Println("\nDay of the week:")
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}
}
