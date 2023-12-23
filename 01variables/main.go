package main

import "fmt"

const LoginValue int = 45

//this is a public variable and its first letter is always written caps

func main() {
	var username string = "arpit"
	fmt.Println(username)
	fmt.Printf("variable type is %T \n", username)
	//%T is placeholder

	var flag bool = true
	fmt.Println(flag)
	fmt.Printf("variable type is %T \n", flag)

	var value int = -24
	fmt.Println(value)
	fmt.Printf("variable type is %T \n", value)

	// int is an alias for int64 or int 32
	// there are variables uint8,uint 16 .. so on  and int8 , int 16 , ....so on
	// there are float32 and float 64

	// default value and some aliases
	var valueOfInt bool
	fmt.Println(valueOfInt)
	fmt.Printf("variable type is %T \n", valueOfInt)
	//so my conclusion is that default values of any int type always 0 ,
	// and of string is empty string and bool has false

	//implicit type
	var hello = "dig"
	fmt.Println(hello)
	// hello=3 this won't work as lexer has already made type hello variable as string

	// := , this walrous operator also used to assign variables , only allowed within func and method
	// not gloabally
	noOftests := 50
	fmt.Println(noOftests)
}
