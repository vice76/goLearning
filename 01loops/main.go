package main

import "fmt"

func main() {
	fmt.Println("loops")

	day := []string{"x", "y", "z", "a", "b", "c"}

	//no kind of like ++d in go
	for d := 0; d < len(day); d++ {
		fmt.Printf("Value at index %v , is %+v\n", d, day[d])
		// fmt.Println(day[d])
	}

	//different type of loop

	for i := range day {
		fmt.Println(day[i])
	}

	for index, day := range day {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	//break , continue and goto

	rogueValue := 1

	for rogueValue <= 10 {
		// if rogueValue == 5 {
		// 	break
		// }
		if rogueValue == 2 {
			goto escape
		}
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("rogueValue is ", rogueValue)
		rogueValue++
	}

escape:
	fmt.Println("Escaping ....")

}
