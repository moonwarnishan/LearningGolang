package main

import (
	"fmt"
)

func main(){
	days :=[]string{"saturday","sunday","Monday","tuesday","wednesday","thursday"}

	// for i:=0; i<len(days); i++{
	// 	fmt.Println(days[i])
	// }

	// for i:=range days{
	// 	fmt.Println(days[i])
	// }

	// i := 1
	// for i<10{
	// 	fmt.Println(i)
	// 	i++;
	// }

	for index, day :=  range days{
		fmt.Printf("index %v and value %v\n",index, day)
	}

}