package main

import (
	"fmt"
)

func main() {

	max := 4000000
	fib := 0
	last := 0
	temp := 0
	counter := 0

	for fib < max {
		if fib == 0 {
			fib++
			continue
		}
		temp = fib
		fib = fib + last
		last = temp
		if fib < max {
			if fib%2 == 0 {
				fmt.Println(fib)
				counter = counter + fib
			}
		}
	}
	fmt.Println("Answer: ", counter)
}
