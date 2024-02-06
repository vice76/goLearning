package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//go.sum is to verify the hash and check whther you are
// downloading the correct thing or not

// note : all the packages go in cache not
// in ur working directory
func main() {
	fmt.Println("hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/products/{key}", serveHome).Methods("GET")

	//connection to a port
	//use command to build the project : go build .
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}
