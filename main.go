package main

import "fmt"

func main() {
	// Variables
	var a int = 10
	b := 20

	// Conditional Statements
	if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is greater than or equal to b")
	}

	// Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Switch Case
	day := 2
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Another day")
	}
}
