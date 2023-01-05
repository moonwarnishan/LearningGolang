package main

import (
	"fmt"
)

func main() {
	Nishan := User{"Nishan", "nishan.rumc@gmail.com", 23, true}
	fmt.Println(Nishan)
	fmt.Printf("%+v", Nishan)
	fmt.Printf("name is %v, status is %v\n", Nishan.Name, Nishan.Status)

	Nishan.status()
	Nishan.EmailChange()
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) status(){
	fmt.Println("user status is ", u.Status)
}

func (u User) EmailChange(){
	u.Email = "CHanges@gmail.com"
	fmt.Println("user email is: ",u.Email)
}