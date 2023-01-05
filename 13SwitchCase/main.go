package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	case 6:
		fmt.Println("six")
	default:
		fmt.Println("Something went wrong")
	}
}
