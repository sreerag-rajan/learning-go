package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://lco.dev"

func main() {
	fmt.Println("WEB Request through GOLANG")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)
	// fmt.Println("Response:", response)

	defer response.Body.Close() //!IMPORTANT  Caller's responsibility

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
