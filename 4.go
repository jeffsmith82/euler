package main

import (
	"fmt"
	"strconv"
)

func main() {

	value := 1000
	max := 0
	left := 0
	right := 0
	for i := 0; i < value; i++ {
		for j := 0; j < value; j++ {
			if i*j > max && palindrone(i*j) {
				max = i * j
				left = i
				right = j
			}
		}
	}
	fmt.Printf("Max:%d , %d x %d\n", max, left, right)
}

func palindrone(input int) bool {
	s := strconv.Itoa(input)

	//Odd number of digits so cant be a palindrone exit early
	if len(s)%2 != 0 {
		return false
	}

	for i := 0; len(s)/2 >= i; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
