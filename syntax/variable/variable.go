package variable

import "fmt"

func Variable() {

	//declare a variable

	var name string = "John Doe"
	var age int = 25
	var isMarried bool = false
	var height float32 = 5.8

	println("Name:", name)
	println("Age:", age)
	println("Married:", isMarried)
	println("Height:", height)

	//declare multiple variables

	var a, b, c, d int = 1, 2, 3, 4

	fmt.Println(a, b, c, d)
}