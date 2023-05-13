package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println("Current Day", days[d])
	// }

	// for i := range days {    //i is returning index and not value
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v and Value is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("Index is something and Value is %v\n", day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			rougueValue++
			continue
		}
		if rougueValue == 4 {
			goto temp
		}
		if rougueValue == 5 {
			break
		}
		fmt.Println("Value is", rougueValue)
		rougueValue++
	}

temp:
	fmt.Println("PRINTING TEMP")
}
