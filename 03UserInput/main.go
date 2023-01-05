package main

import (
	"os"
	"bufio"
	"fmt"
)

func main(){
	welcome := "Welcome to user Input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza")
	// comma ok || error ok
	input,_ := reader.ReadString('\n')
	fmt.Println("Thank You for Input , ",input);
	fmt.Printf("Type of %T", input)

}