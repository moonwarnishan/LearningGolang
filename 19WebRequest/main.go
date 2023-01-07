package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const Url = "https://www.google.com/"
func main(){
	response , err := http.Get(Url);
	ErrorHandle(err)
	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)
	ErrorHandle(err)
	fmt.Println(string(databyte))
}

func ErrorHandle(err error){
	if err!=nil{
		panic(err)
	}
}