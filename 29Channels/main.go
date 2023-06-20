package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	//RECIEVE ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)

		wg.Done()
	}(myCh, wg)

	//SEMD ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
	// myCh <- 5
	// fmt.Println(<-myCh)
}
