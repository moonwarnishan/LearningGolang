package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello array")
	var fruitlist [4]string
	fruitlist[0]="Apple"
	fruitlist[1]="Orange"
	fruitlist[2]="jack fruit"
	fruitlist[3]="tomato"
	fmt.Println("array list : " , fruitlist)
	fmt.Println(len(fruitlist))

	var vegList =[4]string{"tomato","bringal","leaf","red leaf"}
	fmt.Println(vegList)
}