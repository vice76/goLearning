package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welc := "WELCOME USER"
	fmt.Println(welc)

	//here we are using two pkgs that are bufio and os
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the value")

	//comma of || error ok --syntax
	//no try catch
	input, _ := reader.ReadString('\n')
	fmt.Println("output is ", input)

	// type of input is string , good to know , even when user gives a int value .
	fmt.Printf("type of input variable is %T", input)
}
