package language

import "fmt"

/* interfaces are very useful when different structs
 *	use a same method in common. Also for keeping an standard
 */

type Figures2D interface {
	area() float64
}

type square struct {
	base float64
}

type triangle struct {
	base   float64
	height float64
}

func (sq square) area() float64 {
	return sq.base * sq.base
}

// The func area() can be repeated only cause they are methods reffering to an interface
func (trian triangle) area() float64 {
	return trian.base * trian.height
}

func interfaces() {
	mySquare := square{base: 2}
	myTriangle := triangle{base: 2, height: 3}

	fmt.Println("Area square: ", mySquare.area())
	fmt.Println("Area triangle: ", myTriangle.area())

	// Interfaces slice
	myInterfaceSlice := []interface{}{"Nico", 190, "Boca", true}
	fmt.Println(myInterfaceSlice...)
}
