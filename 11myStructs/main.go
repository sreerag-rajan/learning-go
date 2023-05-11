package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//No inheritiance, or super or parent in Golang
	sreerag := User{"Sreerag", "sr@dev.com", true, 28}

	fmt.Println("Sreerag", sreerag)
	fmt.Printf("Sreerag details are %+v\n", sreerag)
	fmt.Printf("Name is %v and age is %v\n", sreerag.Name, sreerag.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
