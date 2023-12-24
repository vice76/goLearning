package main

import "fmt"

func main() {
	fmt.Println("welcome to ifelse classes")

	loginCount := 23

	var ans string
	// these paranthesis should start from the same line otherwise lexer throws a error

	if loginCount < 10 {
		ans = "regular user"
	} else {
		ans = "chal ja ..."
	}
	fmt.Println(ans)

	// some kind of weird syntax
	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than 3")
	}

}
