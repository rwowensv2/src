package main

import "fmt"

// https://golangbot.com/structs/
type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {

	// create examples
	emp1 := Employee{
		firstName: "Rob",
		lastName:  "Pike",
		age:       62,
		salary:    300000,
	}

	emp2 := Employee{"Ben", "Williams", 48, 150000}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
}
