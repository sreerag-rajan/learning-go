package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//No inheritiance, or super or parent in Golang
	sreerag := User{"Sreerag", "sr@dev.com", true, 28}

	fmt.Println("Sreerag", sreerag)
	fmt.Printf("Sreerag details are %+v\n", sreerag)
	fmt.Printf("Name is %v and age is %v\n", sreerag.Name, sreerag.Age)

	sreerag.GetStatus()
	sreerag.SetEmail()
	fmt.Println(sreerag.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}

func (u User) SetEmail() {
	u.Email = "test@email.com"
	fmt.Println("Email of user now is", u.Email)
}
