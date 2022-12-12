package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

 	MyName := "Nols"
	fmt.Printf("Hello I'm %v\n What's your name?\n", MyName)
	
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello %v\n,",name)

}
