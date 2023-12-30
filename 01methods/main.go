package main

import "fmt"

func main() {
	fmt.Println("hello this method class")
	//when functions go into the classes they are called methods

	arpit := User{"arpit", "arpit@web.dev", true, 23}
	// fmt.Println(arpit)
	// fmt.Printf("arpit details are : %+v\n", arpit)
	// fmt.Printf("Name is %v and email is %v", arpit.Name, arpit.Email)
	arpit.GetStatus()
	// so always remember func pass a copy of object
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	//age int
	//it won' be exportable
}

func (u User) GetStatus() {
	fmt.Printf("Name is %v and status is %v \n", u.Name, u.Status)
}
