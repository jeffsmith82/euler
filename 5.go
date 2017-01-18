package main

import (
	"fmt"
)

func main() {

	for i := 21; ; i++ {

		if evenlydiv(i) {
			fmt.Println(i)
			return
		}
	}
}

func evenlydiv(i int) bool {

	for j := 1; j <= 20; j++ {
		if i%j != 0 {
			return false
		}
	}
	return true
}
