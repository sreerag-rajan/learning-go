package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition")

	wg := &sync.WaitGroup{}
	// mut := &sync.Mutex{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("One Routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Second Routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Third Routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Third Routine")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}
