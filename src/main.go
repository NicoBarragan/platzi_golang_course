package main

// if we don't use the library anymore, it will automatically remove
import (
	"fmt"
)

func main() {

	// constants
	const pi float64 = 3.14
	const pi2 = 3.1415

	// Println is a print with a line jump
	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Printf
	name := "Platzi"
	courses := 500
	fmt.Printf("%s has more than %d courses\n", name, courses) //%s: string %v: int
	// or with %v when we don't know the type (avoid if possible)
	fmt.Printf("%v has more than %v courses\n", name, courses) //%v: no type spec

	// Sprintf -> generates a msg but it saves on a variable and it doesnt print on console
	message := fmt.Sprintf("%s has more than %d courses", name, courses)
	fmt.Println(message)

	// int variables
	// for creating and assigning a value to a new var
	// we use ':='. For redeclaring an existent one we use only '='
	base := 12

	// second way
	var height int = 14

	// third way
	var area int

	// if there are unused variables, it won't compile

	fmt.Println(base, height, area)

	// Zero values
	// Values not assigned are not null, they take a default value
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
	// output:
	// 0 0  false

	// For getting the type of a value:
	/* import (
	 * 	"reflect"
	 *   )
	 * fmt.Println(reflect.TypeOf(myValue))
	 */
	// or with the '%T' on format print
	myValue := 20
	fmt.Printf("The type of my value is %T\n", myValue)
}
