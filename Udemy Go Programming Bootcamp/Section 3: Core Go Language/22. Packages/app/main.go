package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello", stuff.Name)
	intArr := []int{2,3,5,7,11}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println("String array :",strArr)
	fmt.Println("Type of String Array :",reflect.TypeOf(strArr))
}