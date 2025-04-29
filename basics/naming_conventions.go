package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func NamingConventionsMain() {
	// PascalCase
	// Eg. CalculateArea, UserInfo, NewHTTPRequest
	// Structs, interface, enums



	// snake_case
	// Eg. user_id, first_name, http_request
	// file names, variable name constants

	// UPPERCASE
	// Eg. PI, AGE
	// Use case is Constants

	// mixedCase
	// Eg. userID, firstName, httpRequest, isValid
	// Use case is variable names

	const MAXRETRIES = 3
	var employeeID = 1001
	fmt.Println(employeeID)

}

