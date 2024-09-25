package array

import "fmt"

func Array() { 

	//decalre an array with 5 elements and assign values
	var arr1 = [5]int{1, 2, 3, 4, 5}

	//declare an array with 5 elements and assign values
	arr2 := [...]int{1, 2, 3, 4, 5}


	fmt.Println(arr1)
	fmt.Println(arr2)
}