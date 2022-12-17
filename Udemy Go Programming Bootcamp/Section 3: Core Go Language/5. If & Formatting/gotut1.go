package main

import (
	"fmt"
	// "reflect"
	// "strconv"
)

var pl = fmt.Println

func main() {
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("This is an important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("This is an important Birthday")
	} else if (iAge >= 65)  {
		pl("This is an important Birthday")
	} else {
		pl("Not an important Birthday")
	}

	pl("!true = ", !true)


	// --------- Formatted Print ---------
	// Go has its own version of C's printf
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value

	fmt.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A', 3.14, true, 1, 1)
	fmt.Printf("%9f\n", -3.14)
	fmt.Printf("%.2f\n", 4.141592)
	fmt.Printf("%9.f\n", 5.141592)

	sp1 := fmt.Sprintf("%9.f\n", 6.141592)
	pl(sp1)
}
