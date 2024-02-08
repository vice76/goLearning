package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vice76/mongoapi/router"
)

func main() {
	fmt.Println("mongodbapi")
	fmt.Println("Server is getting started...")
	r := router.Router()

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening at port 4000...")

}
