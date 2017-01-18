package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	total := big.NewInt(0)
	for i := 1; i < 2000000; i++ {
		if isPrime(i) {
			//total = total + i
			total.Add(total, big.NewInt(int64(i)))
			fmt.Println(total)
		}
	}
	fmt.Println(total)
}

func isPrime(k int) bool {

	for i := 2; i <= int(math.Floor(math.Sqrt(float64(k)))); i++ {
		if k%i == 0 {
			return false
		}
	}
	return k > 1
}
