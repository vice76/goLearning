package main

import "fmt"

func main() {
	fmt.Println("Maps in golang ")

	//syntax
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"
	fmt.Println("map values are as follows ", languages)

	fmt.Println("js value", languages["js"])

	// for deletion
	delete(languages, "js")
	fmt.Println("map values are as follows ", languages)

	//loops in map

	for key, value := range languages {
		fmt.Printf("for key %v , value is %v \n", key, value)
	}
}
