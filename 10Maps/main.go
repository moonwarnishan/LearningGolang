package main

import (
	"fmt"
)

func main(){
	
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"
	fmt.Println(languages)
	fmt.Println(languages["py"])

	delete(languages,"py")
	fmt.Println(languages)

	for key,value := range languages{
		fmt.Printf("key is %v and value is %v",key,value)
	}
}