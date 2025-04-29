package basics

import (
	"fmt"
	"math"
)

func ArithmeticOperatorsMain() {

	// variable declaration

	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Modulus:", result)


	const p float64 = 22.0 / 7.0
	fmt.Println("Value of p:", p)



	// overflow with signed integers

	var maxInt int64 = 9223372036854775807 // max value that int64 can hold
	fmt.Println("Max int64:", maxInt)

	maxInt += 1
	fmt.Println("Max int64:", maxInt)

	// overflow with unsigned integers

	var uMaxUint uint64 = 18446744073709551615 // max value that uint64 can hold
	fmt.Println("Max uint64:", uMaxUint)

	uMaxUint += 1
	fmt.Println("Max uint64:", uMaxUint)
	
	
	var smallFloat float64 = 1.0e-323
	fmt.Println("Small float:", smallFloat)

	smallFloat /= math.MaxFloat64
	fmt.Println("Small float:", smallFloat)
	
	
	
	
	
}

