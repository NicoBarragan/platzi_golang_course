package main

import (
	"fmt"
	"log"
	"strconv"
)

// task1 (par impar)
func numIsOdd(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

// task 2
func login(userName, pass string) bool {
	return userName == "user" && pass == "password"
}

func main() {
	valor1 := 2

	if valor1 == 1 {
		fmt.Println("Is 1")
	} else {
		fmt.Println("Is not 1")
	}

	// convert str to int
	value, err := strconv.Atoi("53")
	// nil == null
	if err != nil {
		// this is for logging errors
		log.Fatal(err)
	}

	fmt.Println("value", value)

	// task 1
	num1, num2 := 2, 3
	odd1 := numIsOdd(num1)
	odd2 := numIsOdd(num2)

	if !odd1 {
		fmt.Printf("num1: %d is even\n", num1)
	} else {
		fmt.Printf("num1: %d is odd\n", num1)

	}

	if !odd2 {
		fmt.Printf("num2: %d is even\n", num2)
	} else {
		fmt.Printf("num2: %d is odd\n", num2)

	}

	// task 2
	user1, pass1 := "user", "password"
	user2, pass2 := "user2", "password2"

	login1 := login(user1, pass1)
	if login1 {
		fmt.Println("login1 is correct")
	} else {
		fmt.Println("login1 is incorrect")
	}

	login2 := login(user2, pass2)
	if login2 {
		fmt.Println("login2 is correct")
	} else {
		fmt.Println("login2 is incorrect")
	}

	// Exec an error
	value, err = strconv.Atoi("bad")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value", value)
	// output of this:
	// 2022/09/20 12:37:09 strconv.Atoi: parsing "bad": invalid syntax exit status

}
