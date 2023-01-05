package main

import (
	"fmt"
)

func main() {
	number := 7
	var result string
	if number > 10 {
		result = "good"
	} else if number < 10 {
		result = "bad"
	} else {
		result = "average"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("Odd")
	}

}
