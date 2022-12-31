package main

import (
	//	"bufio"
	"fmt"
	//	"time"
	//	"unicode/utf8"
	//	"log"
	//	"math/rand"
	//	"os"
	//	"strconv"
	//	"strings"
	// "time"
)

var pl = fmt.Println

func getSum(x, y int) int {
	return (x + y)
}

func getQuotient(x, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		return (x / y), nil
	}
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func getSum2(nums ...int) int {
  sum := 0
  for _, num := range nums {
    sum += num
  }
  return sum
} 

func getArraySum(arr []int) int {
  sum := 0
  for _, num := range arr {
    sum += num
  }
  return sum
} 

func changeVal(f3 int) int {
  // f3 is local to this function
  f3 += 1
  return f3
}

func main() {
	//func funcName(parameters) returnType {BODY}
	pl("getSum(5, 5) : ", getSum(5, 5))

	ans, err := getQuotient(5, 3)
	if err == nil {
		fmt.Printf("%f / %f = %f\n", 5.0, 3.0, ans)
	} else {
		pl(err)
	}

  f1, f2 := getTwo(5)
  fmt.Printf(" f1 end f2 : %d and %d\n", f1, f2)

  unknownSum := getSum2(1,2,3,4,5)
  pl("Unknown Sum : ", unknownSum)

  vArr := []int{1,2,3,4,5,6}
  pl("Array Sum :", getArraySum(vArr))

  // Now we pass by value and not by reference
  f3 := 5
  pl("f3 before func :", f3)
  changeVal(f3)
  pl("f3 after func :", f3)
}
