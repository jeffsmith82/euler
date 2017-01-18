package main

import (
	"fmt"
	"math"
)

func main() {

	total := 0
	results := SieveOfEratosthenes(10)
	//fmt.Println(results)
	for k, v := range results {
		if v == false {
			total = total + k + 1
		}
	}
	fmt.Println(results)
	fmt.Println(total)

}

func SieveOfEratosthenes(value int) []bool {
	f := make([]bool, value)
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if f[i] == false {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < value; i++ {
		if f[i] == false {
			//fmt.Printf("%v ", i)
		}
	}
	return f
}
