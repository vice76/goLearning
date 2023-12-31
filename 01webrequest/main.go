package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("this is webrequest class")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ressponse is of type : %T\n", response)
	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataBytes))
}
