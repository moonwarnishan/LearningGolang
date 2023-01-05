package main

import (
	"fmt"
)

func main(){
	//resutl := add(4,5)
	//fmt.Println(resutl)
	result , message :=  proAdder(1,3,4,6);
	fmt.Println(result)
	fmt.Println(message)
}

func add(valOne int, valTwo int) int {
	return valOne+valTwo
}

func proAdder(values ...int) (int,string){
	total :=0
	for _,val := range values {
		total+=val
	}
	return total ,"total sum"
}