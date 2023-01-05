package main

import (
	"fmt"
)

func main(){
	Nishan := User{"Nishan","nishan.rumc@gmail.com",23,true}
	fmt.Println(Nishan)
	fmt.Printf("%+v",Nishan)
	fmt.Printf("name is %v, status is %v",Nishan.Name, Nishan.Status)
}

type User struct{
	Name string
	Email string
	Age int
	Status bool
}