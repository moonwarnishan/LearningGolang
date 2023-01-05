package main

import (
	"fmt"
)

func main(){
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("location is: ",ptr)
	fmt.Println("value is ", *ptr+3)
	*ptr = *ptr + *ptr
	fmt.Println(*ptr)
}