package main

import (
	"fmt"
)

func main() {

	max := 0
	term := 0
	counter := 1

	for i := 1; i < 13; i++ {
		counter = 1
		for j := i; j != 1; {
			if j%2 == 0 {
				j = j / 2
			} else {
				j = (j * 3) + 1
			}
			counter++
		}
		fmt.Println(i, counter)
		if max < counter {
			max = counter
			term = i
		}
	}
	fmt.Println("Term: ", term)
	fmt.Println("Counter: ", counter)
}
