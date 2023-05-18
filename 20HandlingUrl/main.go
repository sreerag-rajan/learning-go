package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=fjjd87767"

func main() {
	fmt.Println("Welcome to handling URLs in GOlang")
	fmt.Println(myURL)

	//parsing
	result, err := url.Parse(myURL)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme, "\n", result.Host, "\n", result.Path, "\n", result.RawQuery, "\n", result.Port())

	qParams := result.Query()
	fmt.Printf("Type of qParams are %T\n", qParams)

	fmt.Println(qParams["courseName"])

	for _, value := range qParams {
		fmt.Println("Param is ", value)
	}

	partsOfUrl := &url.URL{ //!ALWAYS REMEMBER THE &
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}
