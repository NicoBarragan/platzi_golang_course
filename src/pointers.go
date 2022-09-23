package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

// This is for making a method on a struct
// func (struct) nameFunc()
func (myPC pc) ping() {
	fmt.Println(myPC.brand, "pong")
}

// By working with pointers is more efficient than with the variables
func (myPC *pc) updateRAM(newRam int) {
	myPC.ram = newRam
}

func pointers() {

	// The compiler saves a memory direction to a variable with a value
	a := 50

	// '&<var>' is for accessing at the memroy location
	b := &a

	// '*<mem-location>' is for accessing at the value of a location
	c := *b

	fmt.Println("a: ", a) // 50
	fmt.Println("b: ", b) // 0x14000126008
	fmt.Println("c: ", c) // 50

	// If we change a value directly from the memory location, the variables pointing there will change too
	*b = 100
	fmt.Println("a: ", a) // 100

	myPC := pc{ram: 16, disk: 512, brand: "Apple"}
	fmt.Println("myPC", myPC)

	myPC.ping() // Apple Pong

	// myPC.ram = 20 // inefficient --> updateRAM() does it more optimal
	fmt.Println("First RAM: ", myPC.ram)
	myPC.updateRAM(32)

	fmt.Println("Second RAM: ", myPC.ram)
	myPC.updateRAM(128)

	fmt.Println("Third RAM: ", myPC.ram)

}
