package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	MyName := "Nols"
	fmt.Printf("Hello I'm %v. What's your name?:", MyName)

	var name string
	var age uint
	fmt.Scan(&name)
	fmt.Printf("Hello %v. How old are you?:", name)
	fmt.Scan(&age)

	if age > 10 {
		if age > 60 {
			fmt.Printf("\nFuck!, %v is the age of a Tabotie tree!!", age)
		}
		fmt.Println("\nLet's play!")

	} else {
		fmt.Println("You are to young to play this game, sorry!")
		return
	}

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer)

	fmt.Println(answer)
}
