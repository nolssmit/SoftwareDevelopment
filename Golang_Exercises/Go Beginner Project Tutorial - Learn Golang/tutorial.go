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

	score := 0
	num_questions := 2

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3090" || answer+" "+answer2 == "rtx 3090"{
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9  3900x have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
	  fmt.Println("Correct!")
	  score += 1
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v out of %v.\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100

	fmt.Printf("You scored: %v%%.\n", percent)
}
