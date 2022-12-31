package main

import (
	//	"bufio"
	"fmt"
	"unicode/utf8"
	//	"log"
	//	"math/rand"
	//	"os"
	//	"strconv"
	"strings"
	// "time"
)

var pl = fmt.Println

func main() {
	sV1 := "A word"
// Escape Sequences : \n  \t  \"  \\
	aReplacer := strings.NewReplacer("A", "Another")
	sV2 := aReplacer.Replace(sV1)
	pl(sV2)

	pl("Length :", len(sV2))
	pl("Contains Another :", strings.Contains(sV2, "Another"))
	pl("o index :", strings.Index(sV2, "o"))
	pl("Replace :", strings.Replace(sV2, "o", "0", -1))

	sV3 := "\nSome Words\n"
	sV3 = strings.TrimSpace(sV3)
	pl(sV3)

	pl("Split :", strings.Split("a-b-c-d", "-"))
	pl("Lower :", strings.ToLower(sV2))
	pl("Upper :", strings.ToUpper(sV2))

	pl("Prefix :", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix :", strings.HasSuffix("tacocat", "cat"))

	// Runes
	rStr := "abcdefg"
	pl("Rune Count :", utf8.RuneCountInString((rStr)))

	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}

}
