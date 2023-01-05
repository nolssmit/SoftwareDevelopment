package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
	// "log"
	// "os"
)

var pl = fmt.Println

/*
 +----+
 0	  |
/|\	  |
/ \	  |
     ===

Secret Word : M_N___
Incorrect Guesses : A

Guess a Letter : Y

Sorry Your Dead! The word is MONKEY
Yes the Secret Word is MONKEY

Please Enter Only One Letter
Please Enter a Letter
Please Enter a Letter you Haven't Guessed

*/

var hmArr = [7] string {
	" +----+\n" +
	"      |\n" +
    "      |\n" +
    "      |\n" +
	"     ===\n",
	" +----+\n" +
	" o    |\n" +
    "      |\n" +
    "      |\n" +
	"     ===\n",	
	" +----+\n" +
	" o    |\n" +
    " |    |\n" +
    "      |\n" +
	"     ===\n",
	" +----+\n" +
	" o    |\n" +
    "/|    |\n" +
    "      |\n" +
	"     ===\n",	
	" +----+\n" +
	" o    |\n" +
    "/|\\   |\n" +
    "      |\n" +
	"     ===\n",	
	" +----+\n" +
	" o    |\n" +
    "/|\\   |\n" +
    "/     |\n" +
	"     ===\n",	
	" +----+\n" +
	" o    |\n" +
    "/|\\   |\n" +
    "/ \\   |\n" +
	"     ===\n",					  
}

var wordArr = [7] string {
	"JAZZ", "ZIGZAG", "ZILCH", "ZIPPER",
	"ZODIAC", "ZOMBIE", "FLUFF",
}

var randWord string
var guessedLetters string
var correctLetters []string
var wrongGuesses []string

func getRandWord() string {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randWord = wordArr[rand.Intn(7)]
	correctLetters = make([]string, len(randWord))
	return randWord
}

func showBoard() {
	fmt.Println(hmArr[len(wrongGuesses)])
	fmt.Print("Secret Word : ")
	for _, v := range correctLetters {
		if v == "" {
			fmt.Print("_")
		} else {
			fmt.Print(v)
		}
	}
	fmt.Print("\nIncorrect Guesses : ")
	if len(wrongGuesses) > 0 {
		for _, v := range wrongGuesses {
			fmt.Print(v + "")
		}
	}
	fmt.Println()
}

func getUserLetter() string {
	// Setup buffered reader that gets text from the keyboard
  	reader := bufio.NewReader(os.Stdin) 
	// The guess the user makes
	var guess string

	for true {
		fmt.Print("\nGuess a Letter : ")
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)	
		}
		guess = strings.ToUpper(guess)
		guess = strings.TrimSpace(guess)

		var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
		if len(guess) != 1 {
			fmt.Println("Please Enter Only One Letter")
		} else if !IsLetter(guess) {
			fmt.Println("Please a Letter")
		} else if strings.Contains(guessedLetters, guess) {
			fmt.Println("Please Enter a Letter you Haven't Guessed")
		} else {
			return guess
		}
		return guess
	}
	return guess
}




func main() {
	fmt.Println(getRandWord())
 
	// Show Game Board
	showBoard()

	// Get a letter from the user

	// Add to correctLetters

	// A. If they guessed letter in word
	// Add to correctLetter
	// 1. Are there more letters to guess?
	// 2. If no more letters to guess (YOU WIN)

	// B. If they guessed letter not in word
	// 1. Add new letter to guessedLetters, wrongGuesses
	// 2. Check if they died
	

}
