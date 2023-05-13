package main

import "fmt"

func main() {
	fmt.Println("Learning Functions in Golang")

	greeter()

	var result int = adder(2, 3)

	fmt.Println("Sum", result)

	var bigSum int = proAdder(23, 32, 34, 32, 3, 55, 3, 23)

	fmt.Println("BIG SUM", bigSum)

}

func greeter() {
	fmt.Println("Hello!")
}

func adder(a int, b int) int {
	return a + b

}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}

	return total
}
