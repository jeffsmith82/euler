package main

import (
	"fmt"
	"math"
)

func main() {

	var number int64 = 600851475143

	for i := number / 2; i > 2; i-- {
		fmt.Println(i)
		if number%i == 0 {
			fmt.Println("Prime")
			if isPrime(i) {
				fmt.Println(i)
				return
			}
		}
		i = i - 2
	}
}

func isPrime(k int64) bool {

	fmt.Println(k)
	var i int64 = 2
	for ; i <= int64(math.Floor(math.Sqrt(float64(k)))); i++ {
		if k%i == 0 {
			return false
		}
	}
	return k > 1
}
