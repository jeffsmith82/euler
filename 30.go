package main

import (
	"fmt"
	//	"math/big"
	"strconv"
)

func main() {

	length := 4
	stringy := ""

	for i := 0; len(strconv.Itoa(i)) <= length; i++ {
		stringy = strconv.Itoa(i)
		for j := 0; j < length; j++ {
			//EXP
		}
		fmt.Println(i)
	}
}
