package main

import "fmt"

// variable function

func variable() {
	// Declare variables with initial values

	// Method 1: using var
	var name string = "Mohammed"
	var age int = 21
	var isDev bool = true
	// Method 2: type inference
	var city = "Samara"

	// Method 3: short declaration (inside functions only)
	country := "Ethiopia"

	// Print the variables
	fmt.Println("The value of variables are:", name, age, isDev, city, country)

}

// constant function

func constant() {
	// constant declaration
	const pi = 3.14
	const gravity = 9.8
	const Learning string = "Go"

	// Print the constants
	fmt.Println("The value of constants are:", pi, gravity, Learning)

}

// function data types

func dataTypes() {
	// Declare variables of different data types
	// data types: string, int, bool, float64, etc.
	var name string = "Mohammed"
	var age int = 21
	var isDev bool = true
	var city = "Samara"
	country := "Ethiopia"

	// Print the variables
	fmt.Println("The value of variables are:", name, age, isDev, city, country)
}

func main() {
	variable()
	constant()
	dataTypes()
}
