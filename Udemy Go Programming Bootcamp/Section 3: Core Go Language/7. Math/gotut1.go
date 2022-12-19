package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println

func main() {
	pl("5 + 4 = ", 5+4)
	pl("5 - 4 = ", 5-4)
	pl("5 * 4 = ", 5*4)
	pl("5 / 4 = ", 5/4)
	pl("5 % 4 = ", 5%4) // Get remainder after division

	mInt := 1
	mInt += 1  // Increment with one, same as mInt++
	pl("mInt = ", mInt)

	pl("Float Precision = ", 0.7777777777 + 0.77)

	pl("---------------- Random numbers -------------")
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	

}
