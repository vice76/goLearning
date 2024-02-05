package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to JSON video")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS BootCamp", 299, "LearnCodeOnline.in", "abc123", []string{"js", "java"}},
		{"MERN BootCamp", 300, "LearnCodeOnline.in", "abc456", []string{"react", "mongo"}},
		{"Net BootCamp", 400, "LearnCodeOnline.in", "abc789", nil},
	}

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
