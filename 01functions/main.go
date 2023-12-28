package main

import "fmt"

func main() {
	fmt.Println("welcome to class of function")
	greeter()
	funcExpression := (func() {
		fmt.Println("hello")
	})
	funcExpression()
	result := adder(10, 12)
	fmt.Println(result)
	// var values=make(int, {1,2,3,4,5,6})
	proRes := proAdder(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(proRes)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func greeter() {
	fmt.Println("namaste")

}
