//last in first out
package main

import (
	"fmt"
)

func main(){
	defer fmt.Println("Hello1")
	fmt.Println("Hello")
	Mydefer()
}

func Mydefer(){
	for i:=0; i<5; i++{
		defer fmt.Println(i)
	}
}