package main

import "fmt"

// Structs are the classes in Go
type car struct {
	brand   string
	year    int
	color   string
	seating int
}

func main() {
	// create an instance of a struct
	myCar := car{brand: "Ford", year: 2020, seating: 4}
	fmt.Println("myCar: ", myCar)
	// output:
	// myCar:  {Ford 2020  4}

	// Other way to instancing a struct
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println("otherCar: ", otherCar)
	// Output:
	// otherCar:  {Ferrari 0  0}
}
