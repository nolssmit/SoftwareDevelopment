package main

import (
//	"bufio"
	"fmt"
	// "log"
	// "os"
	// "strconv"
	// "strings"

)

var pl = fmt.Println
func main() {
  fmt.Print("Hello World!\n")
// for intialization; condition; postStatement; {BODY}

  pl("-----------increment-----------")
	for x := 1; x <= 5; x++ {
		pl(x)
	}
	pl("-----------decrement-----------")
	for x := 5; x >= 1; x-- {
		pl(x)
	}	
//	pl("x = ",x)  error, because x cease to exist when lopp completes

pl("--------Simulate a while loop--------")
	fx := 0
	for fx < 5 {
		pl(fx)
		fx++
	} 	
	pl("---------Cycle thru an array of numbers-------")
	aNums := [] int{1, 2, 3}
		for _, num :=range aNums {
			pl(num)
		}

		pl("------Conditionally cycle thru an array of numbers-----")
		xVal := 1
		for true {
			if xVal == 5 {
				break
			}
			pl(xVal)
			xVal++
		}


	}


