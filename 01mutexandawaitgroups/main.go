package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	fmt.Println("Race Condition ")

	var score = []int{0}

	//this leads to race condition
	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("one R ")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("two R ")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()

	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("three R ")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()

	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
