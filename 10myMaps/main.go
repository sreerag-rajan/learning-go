package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string)

	languages["javascript"] = "javascript"
	languages["RB"] = "ruby"
	languages["py"] = "python"

	fmt.Println("languages", languages)

	fmt.Println("py", languages["py"])

	delete(languages, "py")

	for key, value := range languages {
		fmt.Printf("For Key %v, value is %v\n", key, value)
	}
}
