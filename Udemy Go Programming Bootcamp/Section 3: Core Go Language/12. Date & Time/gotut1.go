package main

import (
	//	"bufio"
	"fmt"
	"time"
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
  now := time.Now()
  pl("------- Here in South Africa--------")
  pl(now.Year(), now.Month(), now.Day())
  pl(now.Hour(), now.Minute(), now.Second())

  pl("---------In New York---------")
  loc, err := time.LoadLocation("America/New_York")
  if err != nil {
	pl(err)
  }

  fmt.Printf("Time in New York %s\n", now.In(loc))

  pl("--------Time Standards---------")
  locEST, _ := time.LoadLocation("EST")
  locUTC, _ := time.LoadLocation("UTC")
  locMST, _ := time.LoadLocation("MST")
  fmt.Printf("EST : %s\n", now.In(locEST))
  fmt.Printf("UTC : %s\n", now.In(locUTC))
  fmt.Printf("MST : %s\n", now.In(locMST))

  pl("-------Time Calculations-------")
  birtdate := time.Date(1951, time.June, 2, 12, 30, 10, 0, time.Local)
  pl("Birthday : ", birtdate)
  diff := now.Sub(birtdate)
  fmt.Printf("Days alive : %d days\n", int(diff.Hours()/24))
  fmt.Printf("Hours alive : %d hours\n", int(diff.Hours()))
}
