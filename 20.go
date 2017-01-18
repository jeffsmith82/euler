package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {

	factoral := 100
	//total := 1
	total := big.NewInt(1)
	for i := factoral - 1; i > 0; i-- {
		//total = total * (factoral - i)
		total.Mul(total, big.NewInt(int64(factoral-i)))
		//		fmt.Println(i, total)
	}
	fmt.Println(total)
	totalString := total.String()
	smallTotal := 0
	for j := 0; j < len(totalString); j++ {
		s, _ := strconv.Atoi(string(totalString[j]))
		smallTotal = smallTotal + s
	}
	fmt.Println(smallTotal)
}
