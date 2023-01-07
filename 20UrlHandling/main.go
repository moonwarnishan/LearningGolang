package main

import (
	"fmt"
	"net/url"
)

const URL string = "https://hello.com:3000/learn?username=nishan&password=hello"

func main(){

	result,err := url.Parse(URL)
	ErrorHandle(err)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	query := result.Query()
	fmt.Println(query)
	fmt.Println(query["username"])

	partsOfUrl := &url.URL{
		Scheme:"https",
		Host:"google.com",
		Path:"/google",
		RawPath:"username=nishan",

	}

	anotherUrl := partsOfUrl.String()

 fmt.Println(anotherUrl)
	

}

func ErrorHandle(err error){
	if err!=nil{
		panic(err)
	}
}