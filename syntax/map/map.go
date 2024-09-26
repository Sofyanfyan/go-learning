package maps

import "fmt"

// Map is a collection of key-value pairs. Maps are also known as hash tables, dictionaries, or associative arrays in other programming languages.
func Map() {
	
	// Creating a map
	firtsMap := map[string]any{"name": "Achmad Sofyan Pratama", "age": 24, "city": "Gresik"}

	// Accessing the elements of a map
	firtsMap["height"] = 175
	firtsMap["weight"] = 74

	//checking value of a key
	val1, ok1 := firtsMap["name"]
	val2, ok2 := firtsMap["phone"]

	fmt.Printf("val1: %v, ok1: %v\n", val1, ok1)
	fmt.Printf("val2: %v, ok2: %v\n", val2, ok2)

	// Accessing the elements of a map
	fmt.Printf("firtsMap\t%v\n", firtsMap)
}