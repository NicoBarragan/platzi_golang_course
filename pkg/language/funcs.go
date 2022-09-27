package language

import "fmt"

func normalFunc(message string) {
	fmt.Println(message)
}

// a int, b int
func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

// func     (params) return
func mulTwo(a int) int {
	return a * 2
}

// func         (params) (returns)
func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func funcs() {
	normalFunc("Nico")
	tripleArgument(1, 2, "jaja")

	value := mulTwo(2)
	fmt.Println("Value: ", value)

	value1, value2 := doubleReturn(2)
	fmt.Printf("Value 1: %d | Value2: %d\n", value1, value2)

	// if I don't want to get a var of a return I can use '_'
	_, value3 := doubleReturn(2)
	fmt.Println("Value3: ", value3)

}
