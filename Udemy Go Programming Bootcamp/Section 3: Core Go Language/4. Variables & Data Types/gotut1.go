package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	var vName string = "Nols"
	var v1, v2 = 1.2, 3.5
	v1 = 2.5
	vA := false
	pl(v1, vName, v2)
	pl(reflect.TypeOf(vA))
	pl("---------------------")
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)

	cV3 := "500000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))

	cV5 := 500000
	cV6 := strconv.Itoa(cV5)
	pl(cV6)

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	}

	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9)
}
