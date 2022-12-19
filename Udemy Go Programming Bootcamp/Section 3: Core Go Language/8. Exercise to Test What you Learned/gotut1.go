package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	// "math"
	// "math/rand"
	// "time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//-------------- Read number 1 -------------
	fmt.Print("Enter Number 1 : ")

	num1, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	num1 = strings.TrimSpace(num1)
	iNum1, err := strconv.Atoi(num1)

	if err != nil {
		log.Fatal(err)
	}
	//-------------- Read number 2 -------------
	fmt.Print("Enter Number 2 : ")

	num2, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	num2 = strings.TrimSpace(num2)
	iNum2, err := strconv.Atoi(num2)

	if err != nil {
		log.Fatal(err)
	}

	sum := iNum1 + iNum2
	fmt.Printf("%d + %d = %d\n", iNum1, iNum2, sum)
}
