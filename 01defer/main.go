package main

import "fmt"

func main() {
	fmt.Println("this is defer class")
	justDefering()
	defer fmt.Println("hello")
	fmt.Println("world")
	defer fmt.Println("hello12")
	//line no 7 will execute at last of the function
	//and it follows the principle of LIFO
	// and its scope remians to that func only in which it is used
}

func justDefering() {
	// defer fmt.Println("justDeferring")
	//0 , 1, 2 , 3, 4
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
