package main

import "fmt"

func main() {
	fmt.Println("Welcome to code on pointer")

	// var ptr *int
	// fmt.Println("Default Value of ptr", ptr)

	myNumber := 23

	var ptr = &myNumber //use & when referencing;

	fmt.Println("Value of actual pointer is ", ptr)  //output : 0xc0000bc00 basuically the address
	fmt.Println("Value of actual pointer is ", *ptr) // this gives the value

	*ptr = *ptr * 3
	fmt.Println("MY new value is", myNumber)

}
