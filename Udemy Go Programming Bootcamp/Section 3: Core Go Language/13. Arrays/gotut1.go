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

func main() {
  var arr1 [5]int
  arr1[0] = 1
  arr2 := [5]int{1,2,3,4,5}
  pl("Index 0 of arr2 :",arr2[0])
  pl("Array Length of arr2 :", len(arr2))
  
  //  Iterate thru array elements
  for i :=0; i < len(arr2); i++ {
    pl(arr2[i])
  }

  fmt.Printf("\nArray elements\n")
  for i, v := range arr2 {
    fmt.Printf("%d : %d\n", i, v)
  }

  pl("\n-------Multi dimentional arrays-------\n")  
  arr3 := [2][3]int {
    {1,2,3},
    {4,5,6},
  }

  for  i := 0;  i < 2;  i++ {
    for j := 0; j < 3; j++ {
      pl(arr3[i][j])
    }
  }

  pl("\n--------String to runes (decimal value of character)-------\n")
  aStr1 := "aceg"
  rArr := []rune(aStr1)
  for _, v := range rArr{
    fmt.Printf("Rune Array : %d\n", v)
  }

  pl("--------Byte array to string--------")
  byteArr := []byte{'a','b','c'}
  bStr := string(byteArr[:])
  pl("I'm a string :", bStr)


}
