package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Receive customer data (Their age)
// What is your age :
// Google how to trim whatespace from input
// Age < 5               Too young for school
// Age == 5              Go to Kindergarten
// Age > 5 or Age <= 17  Go to grade GRADE
// Default               Go to college

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("What is your age? :")

	age, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	age = strings.TrimSpace(age)
	iAge, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	} else {	
		grade := iAge - 5
		if iAge < 5 {
			fmt.Println(" Too young for school")
		} else if iAge == 5 {
			fmt.Println("Go to Kindergarten")
		} else if (iAge > 5) && (iAge <= 17) {
			fmt.Printf("Go to grade %d\n", grade)
		} else {
			fmt.Println("Go to college")
		}
	} 
}
