package structure

import "fmt"

func Struct() {

	type Person struct {
		name string
		age  int
	}

	var person Person

	person.name = "John"
	person.age = 25

	// // Struct literal
	// person2 := Person{name: "Jane", age: 30}

	// // Struct pointer
	// person3 := &Person{name: "Doe", age: 35}

	// // Struct pointer
	// person4 := new(Person)
	// person4.name = "Smith"
	// person4.age = 40

	// // Anonymous struct literal
	// anon2 := struct {
	// 	name string
	// 	age  int
	// }{name: "Anon2", age: 50}

	// // Anonymous struct pointer
	// anon3 := &struct {
	// 	name string
	// 	age  int
	// }{name: "Anon3", age: 55}

	// // Anonymous struct pointer
	// anon4 := new(struct {
	// 	name string
	// 	age  int
	// })
	// anon4.name = "Anon4"
	// anon4.age = 60

	fmt.Println(person)
}