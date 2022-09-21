package main

import "fmt"

// Structs are the classes in Go
type car struct {
	brand   string
	year    int
	color   string
	seating int
}

// Stringers: a func that works as interface for printing
// Is efficient than having to write always a dif print statement
// When printing, it will always have this format
func (car car) String() string {
	// sprintf will create a str but won't print on console directly
	return fmt.Sprintf("My car is a %s %s, from %d with %d seats", car.color, car.brand, car.year, car.seating)
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

	fmt.Println(myCar)
	// My car is a  Ford, from 2020 with 4 seats
	fmt.Println(otherCar)
	// My car is a  Ferrari, from 0 with 0 seats

}
