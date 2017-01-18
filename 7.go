package main

import (
	"fmt"
	"math"
)

func main() {

	counter := 0
	for i := 0; ; i++ {

		if isPrime(i) {
			//fmt.Println("Prime:", i)
			counter++
		}
		if counter == 10001 {
			fmt.Println(i)
			return
		}
	}
}

func isPrime(k int) bool {

	for i := 2; i <= int(math.Floor(float64(k)/2)); i++ {
		if k%i == 0 {
			return false
		}
	}
	return k > 1
}
