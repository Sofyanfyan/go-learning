package slices

import "fmt"

func Slices() {
	arr := []int{1, 2, 3, 4, 5}
	slice := arr[1:3]

	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
	fmt.Println("Panjang Slice:", len(slice))
	fmt.Println("Kapasitas Slice:", cap(slice))

	// Memodifikasi elemen slice
	slice[0] = 10
	slice[1] = 20

	// Append elemen seperti "push" pada JavaScript
	slice = append(slice, 30, 40)
	fmt.Println("Slice Setelah Append:", slice)
	fmt.Println("Panjang Slice Setelah Append:", len(slice))
	fmt.Println("Kapasitas Slice Setelah Append:", cap(slice))
}
