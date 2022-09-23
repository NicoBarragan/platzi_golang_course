package main

import "fmt"

func cycles() {

	// For conditional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("")

	// For while (python way)
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// Task: do an decremental loop
	for i := 10; i == 0; i-- {
		fmt.Println(i)
	}

	// For forever (job) it won't end until Ctrl+c
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}

}
