package main

import (
	"fmt"
)

func main() {

	for i := 1; i < 3000000; i++ {
		counter := 0
		for j := 1; j <= i; j++ {
			counter = counter + j
		}
		//fmt.Println(counter)
		// I have my list triangle numbr now i need to work out the number of factors
		factorcount := 0
		for k := 1; k <= counter; k++ {
			if counter%k == 0 {
				factorcount++
			}
		}
		//fmt.Printf("Number: %d, factor count: %d\n", counter, factorcount)
		if factorcount > 500 {
			fmt.Println(counter)
			return
		}
	}
}
