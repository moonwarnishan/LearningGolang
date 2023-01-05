package main

import ("fmt"
		"time")

func main(){
	fmt.Println("Welcome to time")
	timeNow := time.Now()
	fmt.Println(timeNow)
	//tie format change
	fmt.Println(timeNow.Format("01-02-2006 15:04:05 Monday \n"))

	createDate := time.Date(1999, time.May, 19,12,00,00,00,time.UTC)
	fmt.Println(createDate.Format("01-02-2006 12:00:00 Monday"))
}