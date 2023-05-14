package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Go lang")

	content := "This needs to go into a file"

	file, err := os.Create("./myGoFile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length is", length)

	defer file.Close()

	readFile("./myGoFile.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilError(err)

	fmt.Println("Text Data in the file is \n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
