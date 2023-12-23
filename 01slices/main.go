package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")
	//this data structure is used a lot more time
	// they are under the hood of arrays

	//syntax
	var flist = []string{"Apple", "Tomato", "banana"}
	fmt.Printf("Type os fruitList is %T\n", flist)

	flist = append(flist, "mango", "gauva")
	fmt.Println(flist)

	flist = append(flist[1:])
	fmt.Println(flist)

	// using make()

	highScores := make([]int, 4)

	highScores[0] = 15
	highScores[1] = 14
	highScores[2] = 13
	highScores[3] = 12
	//highScores[4]=1;
	// if you use this , it ain't going to work
	// but wait

	highScores = append(highScores, 655, 766, 877)
	//its going to work as this will reallocate the memory and tries to allocae all
	// the values inserted after the slices size
	// saves lot of time and good example of perfomance optimization

	// use sort pkg to sort this
	sort.Ints(highScores)
	fmt.Println(highScores)

	//how to remove a value from slices based on INdex
	var index int = 2
	highScores = append(highScores[:index], highScores[index+1:]...)
	fmt.Println(highScores)

}
