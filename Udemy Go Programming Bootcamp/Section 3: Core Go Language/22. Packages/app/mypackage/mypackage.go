package stuff

import (
	//"errors"
	"strconv"
	//"time"
)
// Variables and functions starting with a 
// capital letter can be accessed from outside
var Name string = "Nols"  

func IntArrToStrArr(intArr []int) []string{
	var strArr [] string
	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}