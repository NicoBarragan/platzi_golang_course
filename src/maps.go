package main

import "fmt"

func main() {
	// the keyword 'make' is usful when declaring variables
	ageMap := make(map[string]int)

	// setting values
	ageMap["Nico"] = 20
	ageMap["Patri"] = 24

	// Iterating over a map w range
	for name, age := range ageMap {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// getting values (ok is a boolean that says id the value exist or not in the map)
	// is useful because if we don't have a value, it will return the zero value of the type
	nicoAge, ok := ageMap["Nico"]
	fmt.Printf("nicoAge: %d, boolean ok: %t\n", nicoAge, ok)

	juanAge := ageMap["Juan"]
	fmt.Println("Juan age (without ok boolean): ", juanAge)

	juanAge, ok = ageMap["Juan"]
	fmt.Printf("Juan age: %d, ok boolean: %t\n", juanAge, ok)

}
