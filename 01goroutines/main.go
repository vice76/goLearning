package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}
var wg sync.WaitGroup //pointer

var mut sync.Mutex //pointer

func main() {
	// go greeter("Hello")
	// greeter("world")
	//not waited so doesn't prints
	website := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
	}

	for _, web := range website {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		//as soon as the thread comes back it prints
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	}

	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
}
