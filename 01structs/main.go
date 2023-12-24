package main

import "fmt"

func main() {
	fmt.Println("welcome to class of structs")

	//no inheritance , no concept of super or parent
	arpit := User{"arpit", "arpit@web.dev", 23, true}
	fmt.Println(arpit)
	fmt.Printf("arpit details are : %+v\n", arpit)
	fmt.Printf("Name is %v and email is %v", arpit.Name, arpit.Email)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
