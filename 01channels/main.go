package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in golang")

	// myCh := make(chan int)
	myCh := make(chan int, 2) //buffer chnnel with no of values a channel can get

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh
		//way to handle closed channel

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

	// myCh <- 5
	// fmt.Println(<-myCh)

	// channels says first listen to me then I will
	//allow you to pass the value
	//sol is go routine
}
