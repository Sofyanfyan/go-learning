package constant

import "fmt"

func Constant() {	


// 	If a variable should have a fixed value that cannot be changed, 
//you can use the const keyword.
// The const keyword declares the variable as "constant", which means 
//that it is unchangeable and read-only.


	const name string = "John Doe";
	fmt.Println(name)
}