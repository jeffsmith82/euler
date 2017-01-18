package main

import (
	"fmt"
	"math"
)

func main() {

	squares := 0
	count := 0

	for i := 1; i <= 100; i++ {
		squares = squares + int(math.Pow(float64(i), 2))
		count = count + i
	}

	//fmt.Println(squares)
	total_squared := int(math.Pow(float64(count), 2))
	fmt.Println(total_squared - squares)

}
