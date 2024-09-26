package interfaces

import "fmt"

type Animal interface {

	Speak() string
}

type Dog struct {
	name string
}

type Cat struct { 
	name string
}

func (d Dog) Speak() string {
	
	return "Woof!"
}

func (c Cat) Speak() string {
	return "Meow!"
}

// Interface is a collection of method signatures that a Type can implement.
func Interface() {

	// poodle := Dog{"Poodle"}
	// fmt.Println(poodle.Speak())

	// var animal Animal
	// animal = poodle
	// fmt.Println(animal.Speak())

	animals := []Animal{Dog{"Poodle"}, Cat{"Persian"}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}