package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("welcome to handling to web requests")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code :", response.StatusCode)
	fmt.Println("Content Length is :", response.ContentLength)

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content is :", string(content))

}
