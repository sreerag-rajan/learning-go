package main

import "fmt"

func main() {
	fmt.Println("If Else in golang")

	loginCount := 20

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "User is hooked"
	} else {
		result = "Exact 10"
	}

	fmt.Println("Result", result)

	if 9%2 == 0 {
		fmt.Println("Number is Event")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less thatn 10")
	}

	// if err != nil {

	// }
}
