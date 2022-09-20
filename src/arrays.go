package main

import "fmt"

func main() {
	// Array (fixed)
	var array [4]int
	array[0] = 1
	array[1] = 2

	fmt.Println("array: ", array)
	// len of the actual array
	fmt.Println("len array: ", len(array))
	// max capacity of the array
	fmt.Println("cap array: ", cap(array))

	// Slice: Variable array
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	/* Slice methods */
	// Slicing: array[from(included): to(excluded)]
	fmt.Println(slice[0])
	fmt.Println(slice[1:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append: If there is not space for new elements, a new array is appended so the capacity doubles
	fmt.Println("Cap before slicing slice: ", cap(slice))
	slice = append(slice, 7, 8)
	fmt.Println("Cap slice on 1st slicing: ", cap(slice))

	// Append with other list (like destructuring)
	newSlice := []int{9, 10}
	slice = append(slice, newSlice...)
	// Both must be or fixed or variables arrays for this method to succeed
	fmt.Println("Cap slice on 2nd slicing: ", cap(slice))
	// like there were enoguh empty values in the slice, the capacity won't change this time

	fmt.Println("Slice: ", slice)

	// The cap and len are diff when an array is sliced
	childSlice := slice[:5]
	fmt.Printf("The len of childSlice is %d | The cap of childSlice is %d\n", len(childSlice), cap(childSlice))
	fmt.Printf("The len of slice is %d | The cap of slice is %d\n", len(slice), cap(slice))

}
