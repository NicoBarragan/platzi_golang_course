package main

import "fmt"

// func that adds a string to a channel
func message(text string, c chan string) {
	c <- text
}

func channels() {
	// make(type, maxCapacity)
	c := make(chan string, 2)

	// Adding values to the channel
	c <- "Msj1"
	c <- "Msj2"

	fmt.Println("c cap & len: ", len(c), cap(c))
	// Output: 2, 2

	fmt.Println("-------")

	// Range & close
	// The good practice is to close a channel when it doesn't receive any other data
	close(c)
	// Output:
	// Msj1
	// Msj2

	// If is full or if is it closed, it won't be able to receive more data
	// c <- "Msj3" --> So this will fail

	fmt.Println("-------")

	// Iterating a channel
	for message := range c {
		fmt.Println(message)
	}

	/* Select */
	// creating 2 dynamic channels
	email1 := make(chan string)
	email2 := make(chan string)

	// exec 2 different Goroutines through 2 diff channels w message() func
	go message("msj1", email1)
	go message("msj2", email2)

	// Here we use select in a for loop for getting which channel will arrive first
	// We iterate through the number of channels we have
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email received from email1", m1)
		case m2 := <-email2:
			fmt.Println("Email received from email2", m2)
		}
	}
}
