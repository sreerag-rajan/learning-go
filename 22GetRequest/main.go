package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Methods in GOLANG")

	// performGetRequest()
	performPostFormRequest()

}

func performGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount", byteCount)
	fmt.Println("", responseString.String())

	fmt.Println(string(content))
}

func performPostRequest() {
	const myUrl string = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"courseName":"Lets go with golang",
			"price": 0,
			"platform": "youtube.com"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Data", string(data))
}

func performPostFormRequest() {
	const myUrl = "http://localhost:8000/post"

	data := url.Values{}

	data.Add("firstName", "Sreerag")
	data.Add("lastName", "rajan")

	response, err := http.PostForm(myUrl, data)

	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
