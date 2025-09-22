package main

import "fmt"

// Struct embedding
type Person struct {
	Name string
}

type Student struct {
	Person // embedding
	School string
}

// Each variable in a struct is called a field.
type User struct {
	ID    int
	Name  string
	Email string
}

// struct with methods

type Rectangle struct {
	Width  float64
	Height float64
}

// Method to calculate area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method with pointer receiver (modifies struct)
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	// Struct embedding example
	student := Student{
		Person: Person{Name: "Mohammed"},
		School: "xyz University",
	}
	fmt.Println(student)

	// Struct with fields example
	user := User{ID: 1, Name: "Mohammed", Email: "mohammed@example.com"}
	fmt.Println(user)

	// Struct with methods example
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area:", rect.Area())
}
