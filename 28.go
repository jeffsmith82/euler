package main

import (
	"fmt"
)

func main() {

	var total int64 = 1
	size := 1001
	counter := 2
	loops := 1
	temp := 0
	for i := 1; loops < (size - 2); i++ {
		for j := 1; j <= 4; j++ {
			fmt.Println(counter*j + i)
			temp = counter*j + i
			total = total + int64(temp)
		}
		counter = counter + 2
		i = temp - 1
		fmt.Println("I:", i)
		loops++
	}
	fmt.Println("Total: ", total)
}
