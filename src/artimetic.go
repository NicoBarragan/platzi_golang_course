package main

import "fmt"

func squareRecArea(base, height float64) float64 {
	return base * height
}

func trapezeArea(trapHeight, trapBaseZero, trapBaseOne float64) float64 {
	return trapHeight / 2 * (trapBaseOne + trapBaseZero)
}

func circleArea(ratio float64) float64 {
	const PI float64 = 3.14
	return PI * (ratio * ratio)
}

func main() {
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

	// Calculate the area of square
	squareBase := 10.0
	squareArea := squareRecArea(squareBase, squareBase)
	fmt.Println("Square area: ", squareArea)

	// Calculate rectangle area
	recBase := 10.0
	recHeight := 5.0

	recArea := squareRecArea(recBase, recHeight)
	fmt.Println("RecArea: ", recArea)

	// Calculate trapeze area
	trapBaseZero := 12.0
	trapBaseOne := 17.0
	trapHeight := 10.0

	trapArea := trapezeArea(trapHeight, trapBaseOne, trapBaseZero)
	fmt.Println("trapArea: ", trapArea)

	// Calculate circle area
	circleDiam := 50
	circleRadio := float64(circleDiam / 2)
	circleArea := circleArea(circleRadio)

	fmt.Println("circleArea: ", circleArea)
}
