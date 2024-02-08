package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Hello")
	greeter("world")
	//not waited so doesn't prints
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		//as soon as the thread comes back it prints
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}
