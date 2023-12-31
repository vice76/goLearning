package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("this is class of files")

	content := "this needs to be in file"
	file, err := os.Create("./mylcofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("length is", length)
	defer file.Close()
	readFile("./mylcofile.txt")
}

func readFile(filename string) {
	//for reading file we have different entity known as ioutils
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("text data inside the file is \n", string(dataByte))

}
