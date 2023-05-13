package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in Golang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is one. You can open or move 1 space")
	case 2:
		fmt.Println("DIce value is 2, You can move 2 spaces")
	case 3:
		fmt.Println("DIce value is 3, You can move 3 spaces")
	case 4:
		fmt.Println("DIce value is 4, You can move 4 spaces")
		fallthrough // this is the opposite of break in JS. it allows for other cases to run as well
	case 5:
		fmt.Println("DIce value is 5, You can move 5 spaces")
	case 6:
		fmt.Println("DIce value is 6, You can either open or move 6 spaces, then you can roll again as you get one more turn")
	default:
		fmt.Println("What was that?")
	}

}
