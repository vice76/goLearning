package main

import "fmt"

func main() {
	//arrays are not used so much in this language
	fmt.Println()

	// if you do not assign any value to any index of array , it will
	// automatically assign a default value to it , its like java

	var list [4]int
	list[0] = 1
	list[1] = 2

	list[3] = 4

	fmt.Println("value of array is", list[2])
	fmt.Printf("type of array %T", list[2])

	// more ways of declaration
	var veg = [3]string{"a", "b", "c"}
	fmt.Println("bs bana k dekh ha hu ", len(veg))
}
