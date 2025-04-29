package basics

import "fmt"

func VariablesMain() {
	var age int
	var name string = "John"
	var name1 = "Jane"

	count := 10
	lastName := "Smith"

	// Default values
	// Numberic Types: 0
	// String Types: ""
	// Boolean Types: false
	// Pointer, slices, maps, channels, functions and structs: nil

	// ----- SCOPE

	fmt.Println(age, name, name1, count, lastName)
}

func printname() {
	firstName := "John"
	fmt.Println(firstName)
}
