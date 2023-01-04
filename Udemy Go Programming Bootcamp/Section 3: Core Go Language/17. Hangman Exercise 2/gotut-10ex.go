package main

import (
//	"bufio"
	"fmt"
//	"log"
//	"os"
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

var hArr = [7] string {
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
var correctLetter []string
var wrongGuesses []string

func main() {
 //reader := bufio.NewReader(os.Stdin)  //read from keyboard

	// Show Game Board

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
