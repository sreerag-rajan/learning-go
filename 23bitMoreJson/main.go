package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"courseName"`
	Price    float32  `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON in golang")
	// EncodeJson()
	DecodeJSON()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Learncode.in", "abc123", []string{"Webdev", "js"}},
		{"MERN Bootcamp", 199, "Learncode.in", "hit123", []string{"fulstack", "js"}},
		{"ANGULAR Bootcamp", 299, "Learncode.in", "bcd123", nil},
	}

	// finalJSON, err := json.Marshal(lcoCourses)
	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	// fmt.Println(finalJSON)
	fmt.Printf("%s\n", finalJSON)
}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
		{
			"courseName": "MERN Bootcamp",
			"price": 199,
			"website": "Learncode.in",
			"tags": ["fulstack","js"]
		}
	`)

	var lcocourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)

	} else {
		fmt.Println("JSON was not valid")

	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
}
