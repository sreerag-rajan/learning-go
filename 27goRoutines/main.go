package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

func main() {
	// go greeter("HELLO")
	// greeter("WORLD")
	websiteList := []string{
		"http://lco.dev",
		"http://go.dev",
		"http://google.com",
		"http://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
