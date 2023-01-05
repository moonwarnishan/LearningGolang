package main

import (
	"fmt"
	"io/ioutil"
	"io"
	"os"
)

func main(){
	content := "This is content file"
	file, err := os.Create("./content.txt")
	if err != nil {
		panic(err)
	}
	io.WriteString(file,content)
	defer file.Close()

	ReadFile("./content.txt")
}

func ReadFile(path string){
	dtabyte,err := ioutil.ReadFile(path)
	if err!=nil{
		panic(err)
	}
	fmt.Println("content is: ", string(dtabyte))
}