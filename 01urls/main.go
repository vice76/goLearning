package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=fghjhgfgh56"

func main() {
	fmt.Println("welcome to handling of urls")
	//parsing of url
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	queris := result.Query()

	fmt.Printf("The type of queris is %T\n", queris)
	//type is url.values // something of a hashmap
	fmt.Println(queris["coursename"])

	//craete a new url

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
