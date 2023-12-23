package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("studying about time")
	presentTime := time.Now()
	fmt.Println("time is now", presentTime.Format("01-02-2006 15:04:05 Monday"))
	//there should be the exact syntax for date otherwise it will give error

	//executables
	//GOOS="windows || linux || darwin" go build ---command
}
