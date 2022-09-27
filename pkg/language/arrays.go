package language

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) string {
	var reverseText string

	// iterate from the final to the start the text
	for i := len(text) - 1; i >= 0; i-- {
		reverseText += strings.ToLower(string(text[i]))
		// Why do we parse text[i] with string()? Because an string is a characters list that compresses into ASCII
		// when we pick only an element, we will get uint8 (ASCII) as type, and we need it to be a string
	}

	// Only for learning purposes:
	fmt.Printf("Type of a character of text[0]: %T\n", text[0])

	if reverseText == strings.ToLower(text) {
		return "Is Palindrome"
	}

	return "Is not Palindrome"
}

func arrays() {
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

	/* Range */
	// The range keyword is used in for loop to iterate over items of an array, slice, channel or map.
	// With array and slices, it returns the index of the item as integer.
	// starts from 0 and ends in len
	for i := range slice {
		fmt.Println("Item", i, "is equal to", slice[i])
	}

	// With maps, it returns the key of the next key-value pair. Range either returns one value or two.
	// define map to iterate:
	capitalsMap := map[string]string{"Argentina": "Buenos Aires", "Uruguay": "Montevideo"}

	//  If only one value is used on the left of a range expression, it is the 1st value in the following table.
	for country, capital := range capitalsMap {
		fmt.Printf("%s is the capital of %s\n", country, capital)
	}

	// Palindrome
	fmt.Printf("\n-------------\nPalindrome:\n")
	fmt.Println(isPalindrome("AMa"))
	fmt.Println(isPalindrome("AMar azul"))

}
