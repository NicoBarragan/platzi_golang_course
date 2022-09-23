package main

import "fmt"

func switchs() {
	// For a var without too many conditions this is the best option
	switch module := 5 % 2; module {
	case 0:
		fmt.Println("Is odd")
	default:
		fmt.Println("Is Even")
	}

	// For multiple conds on a var this is better:
	value := 200

	switch {
	case value > 100:
		fmt.Println("Greater than 100")
	case value < 0:
		fmt.Println("Less than zero")

	default:
		fmt.Println("No condition")

	}
}
