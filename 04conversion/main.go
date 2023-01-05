package main

import (
	"strings"
	"strconv"
	"os"
	"fmt"
	"bufio"
)

func main(){
	fmt.Println("Welcome to our pizza shop")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("give pizza rating: ")
	input , _ := reader.ReadString('\n')
	
	convertToFloat,err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if err!=nil {
		fmt.Println(err)
	} else {
		fmt.Println("After Added number is : ",convertToFloat+1)
	}

	



}
