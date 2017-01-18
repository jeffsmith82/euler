package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {

	total := big.NewInt(1)
	for i := 1; i <= 1000; i++ {

		total.Mul(total, big.NewInt(2))
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
