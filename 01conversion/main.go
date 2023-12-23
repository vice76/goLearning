package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("pls enter restaurant rating")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("rating are as follows , ", input)

	//this will cause error bcoz along with the value entered there will \n also coming to the input
	// so need to trim that too using another pkg that is pkg strings
	inputFloat, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
		//also can use Panic(err) to end the program but we won't do this here.
	} else {
		fmt.Println("new ratings are as followed", inputFloat+1)
	}
	//
	inputFloatWithoutErr, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("new ratings are as followed", inputFloatWithoutErr+1)
	}
}
