package main

// if we don't use the library anymore, it will automatically remove
import (
	"fmt"
)

func main() {

	// constants
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

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
	// or
	/* fmt.Printf(" The type of my value is %T", myValue) */

	// Calculate the area of square
	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("Square area: ", squareArea)

	/* Aritmetic operators */
	// Sum
	x := 10
	y := 50

	result := x + y
	fmt.Println("Sum: ", result)

	// Substraction
	result = y - x // we are reasigning a declared var so the ':' are not necessary
	fmt.Println("Sub: ", result)

	// Multiplication
	result = x * y
	fmt.Println("Mul: ", result)

	// Division
	result = y / x
	fmt.Println("Div: ", result)

	// Module
	result = y % x
	fmt.Println("Mod: ", result) // 0

	result2 := x % y
	fmt.Println("Mod 2: ", result2) // 10

	// Incremental
	x++
	fmt.Println("Incremental: ", x) // 11

	// Decremental
	x--
	fmt.Println("Incremental: ", x) // 10

	// Calculate rectangle area
	recBase := 10
	recHeight := 5

	recArea := recBase * recHeight
	fmt.Println("RecArea: ", recArea)

	// Calculate trapeze area
	trapBaseZero := 12
	trapBaseOne := 17
	trapHeight := 10

	trapArea := trapHeight / 2 * (trapBaseOne + trapBaseZero)
	fmt.Println("trapArea: ", trapArea)

	// Calculate circle area
	circleDiam := 50
	circleRadio := float64(circleDiam / 2)
	const PI float64 = 3.14
	circleArea := circleRadio * circleRadio * PI

	fmt.Println("circleArea: ", circleArea)
}
