package main

import (
	"fmt"
)

func main() {
	counter := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			//fmt.Println(i)
			counter = counter + i
		}
	}
	fmt.Println(counter)
}
