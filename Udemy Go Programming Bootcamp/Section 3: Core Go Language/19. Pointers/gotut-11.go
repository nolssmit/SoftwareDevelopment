package main

import (
	"fmt"
)

var pl = fmt.Println

func changeVal2(myPtr *int) {
	*myPtr = 12
}

func dblArrVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, val := range nums {
		sum += val
	}
	return (sum / numSize)
}

func main() {
	f4 := 10

	pl("f4 :", f4)
	pl("f4 Address :", &f4)
	var f4Ptr *int = &f4
	pl("f4 Address :", f4Ptr)
	pl("f4 Value :", *f4Ptr)
	*f4Ptr = 11
	pl("f4 New Value :", *f4Ptr)

	pl("f4 before calling Function :", f4)
	changeVal2(&f4)
	pl("f4 after calling Function :", f4)

	pArr := [4]int{1,2,3,4}
	pl("Array values before doubling",pArr)
	dblArrVals(&pArr)
	pl("Array values after doubling",pArr)

	iSlice := []float64 {11,13,17}
	fmt.Printf("Average : %.3f\n", getAverage(iSlice...))
}