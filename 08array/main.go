package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitlist [4]string //the lenght of array has to be declared

	fruitlist[0] = "apple"
	fruitlist[1] = "orange"
	fruitlist[3] = "peach"

	fmt.Println("Fruit list is ", fruitlist)
	fmt.Println("Length of fruitlist is ", len(fruitlist)) //this will always be what you have declared. Does not matter even if the list is empty the declared value will be shown

	var veglist = [3]string{"Potato", "Carrot", "Tomato"}
	fmt.Println("Vegy list is", veglist)
	fmt.Println("Length of veglist is ", len(veglist))
}
