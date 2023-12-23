package main

import "fmt"

//about pointers
//so what happens we create variables and tehse are being passed to too many functions
// after being passed funtions , they are some irregularities observed in the value
// of varibles , so programmers introduced  something called pointers
// pointers provide the surety , instead they pass the address of that variable that makes sure there
// are no irregularities

func main() {
	fmt.Println("welcome to become a gaint and know about pointers")
	// so the value of uninitialized pointer is nil
	var ptr *int
	fmt.Println("Value of pointer", ptr)

	myNumber := 23
	var ptr1 = &myNumber

	fmt.Println("Value of pointer is ", ptr1)
	fmt.Println("Value of pointer is ", *ptr1)

	// so here is the use of pointer operation doesn't works on copy of varible
	// but does on the exact address
	*ptr1 = *ptr1 * 2
	fmt.Println("New value is ", myNumber)

}
