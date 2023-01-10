package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
/*	
	----- COMMAND LINE ARGUMENTS -----
	You can pass values to your program
	from the command line
	Create gotut2.go
	go build gotut2.go
	.\gotut2 24 43 12 9 10
	Returns an array with everything
	passed with the name of the app
	in the first index
	Outputs the max number passed in
*/
	fmt.Println(os.Args)
	args := os.Args[1:]
	iArgs := []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}
	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}
	fmt.Println("Max Value :", max)
}