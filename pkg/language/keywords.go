package language

import "fmt"

func keywords() {
	// Defer: Executes a line before the func dies
	// It makes like an offset and executes it after all of the lines in the func
	defer fmt.Println("Hello")
	fmt.Println("World")
	// Output (of only this lines): World Hello
	/* Defer is very useful for example when you have to close a connection to a BD, or read a file in order to finish the process */
	// The good practice is using only 1 defer if the function needs, not more

	// continue & break on loops
	for i := 0; i < 10; i++ {

		if i == 2 {
			fmt.Println("Two")
			continue
		}

		if i == 8 {
			fmt.Println("Eight")
			break
		}

		fmt.Println(i)
	}

}

// Final Output:
// World
// 	0
// 1
// Two
// 3
// 4
// 5
// 6
// 7
// Eight

// Hello
