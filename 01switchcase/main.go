package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("creating a dice game using switch case")

	rand.New(rand.NewSource(time.Now().UnixNano())) //just to generate a random number
	dicenumber := rand.Intn(6) + 1                  // gives a number between 1 -6
	fmt.Println("Value of dice id ", dicenumber)

	//fallthrough is not like a break statement , it will also fall in that cases and gets satisfied
	switch dicenumber {
	case 1:
		fmt.Println("Dice valye is 1 and start the game")
	case 2:
		fmt.Println("You can move 2 steps")
	case 3:
		fmt.Println("You can move 3 steps")
		fallthrough
	case 4:
		fmt.Println("You can move 4 steps")
		fallthrough
	case 5:
		fmt.Println("You can move 5 steps")
	case 6:
		fmt.Println("You can move 6 steps")
	default:
		fmt.Println("wrong value ")
	}
}
