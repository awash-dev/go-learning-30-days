package main

import (
	"fmt"
)

func main() {
	// --- ARRAYS ---
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)
	fmt.Println("First element:", arr[0])

	// --- SLICES ---
	slice := []string{"Go", "Python", "Rust"}
	fmt.Println("\nSlice:", slice)

	// Append
	slice = append(slice, "JavaScript")
	fmt.Println("After append:", slice)

	// Slice operator
	sub := slice[1:3] // from index 1 up to 3 (exclusive)
	fmt.Println("Sub-slice:", sub)

	// Copy slice
	copySlice := make([]string, len(slice))
	copy(copySlice, slice)
	fmt.Println("Copied slice:", copySlice)

	// --- MAPS ---
	langMap := map[string]string{
		"go": "Golang",
		"py": "Python",
		"js": "JavaScript",
	}
	fmt.Println("\nMap:", langMap)
	fmt.Println("Get value for key 'go':", langMap["go"])

	// Add new key
	langMap["rs"] = "Rust"
	fmt.Println("After adding Rust:", langMap)

	// Delete key
	delete(langMap, "py")
	fmt.Println("After deleting Python:", langMap)

	// Loop through map
	fmt.Println("\nLoop through map:")
	for key, value := range langMap {
		fmt.Printf("%s => %s\n", key, value)
	}
}
