package main

import (
	"sort"
	"fmt"
)


func main(){
	var fruitlist = []string{"apple","banana"}
	fmt.Println(fruitlist)
	fruitlist = append(fruitlist,"jak","mango")
	fmt.Println(fruitlist[:3])

	highscore := make([]int,4)
	highscore[0] = 111
	highscore[1] = 12
	highscore[2] = 183
	highscore[3] = 143
	//out of bound case
	//highscore[4] =15

	//working method in slices
	highscore = append(highscore,15,16,17)
	fmt.Println(highscore)
	fmt.Println(len(highscore))
	sort.Ints(highscore)
	fmt.Println(highscore)

	//remove value from slice
	var index int = 2;
	highscore = append(highscore[:index],highscore[index+1:]...)
	fmt.Println(highscore)



}