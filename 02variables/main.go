package main

import "fmt"

func main() {
	var username string = "Sreerag"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type : %T \n", smallVal)

	var smallFLoat float32 = 255.878755776787654567
	fmt.Println(smallFLoat)
	fmt.Printf("variable is of type : %T \n", smallFLoat)

	var unassignedValue string
	fmt.Println(unassignedValue)
	fmt.Printf("variable is of type : %T \n", unassignedValue)
}
