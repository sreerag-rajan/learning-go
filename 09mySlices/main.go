package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to learning Slices")

	var fruitList = []string{"apple", "peach", "tomato"}
	fmt.Printf("Type of fruiltlist is %T\n", fruitList)

	fruitList = append(fruitList, "orange", "banana")

	fmt.Println("NEw fruitlist", fruitList)

	fruitList = append(fruitList[1:])

	fmt.Println("NEw fruitlist", fruitList)

	highscore := make([]int, 4)

	highscore[0] = 234
	highscore[1] = 244
	highscore[2] = 254
	highscore[3] = 274

	fmt.Println("highscores", highscore)

	highscore = append(highscore, 234, 435, 345)

	fmt.Println("Highscores", highscore)

	sort.Ints(highscore)
	fmt.Println("HIGH sorted", highscore)

	//removing values from slice based on index

	var courses = []string{"reactjs", "javascript", "Swift", "python", "ruby"}

	fmt.Println("Courses", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("COURSES", courses)
}
