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
	// var name []datatype
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "0f"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	pl("Slice Size :", len(sl1))

  for i := 0; i < len(sl1); i++ {
    pl(sl1[i])
  }

  for _, x := range sl1{
    pl(x)
  }

  sl2 := []int{22, 12, 2022}
  pl(sl2)

  sArr := [7]int{1, 2, 3, 4, 5, 6, 7}
  sl3 := sArr[0:2]
  pl(sl3)

  sl4 := sArr[:3]
  pl("From first index up to, but not including 3'rd :",sl4)

  sl5 := sArr[3:]
  pl("From 3'rd index till end, and including 3'rd :",sl5)  

   sArr[0] = 10
   pl("sl3, where sl3 := sArr[0:2] :",sl3) 

  sl3 = append(sl3, 12)
  pl("sl3, with 12 appended :",sl3)
  pl("sArr, with sArr[0] = 10, replacing original value :", sArr)

  sl6 := make([]string, 6)
  pl("sl6 :", sl6)
  pl("sl6[0] :", sl6[0])
}
