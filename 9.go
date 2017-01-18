package main

import (
	"fmt"
	"math"
)

func main() {

	for a := 1; a < 500; a++ {
		for b := 1; b < 500; b++ {
			for c := 1; c < 500; c++ {
				if (a + b + c) == 1000 {
					//Check if pythag
					as := int(math.Pow(float64(a), 2))
					bs := int(math.Pow(float64(b), 2))
					cs := int(math.Pow(float64(c), 2))
					if as+bs == cs {
						fmt.Printf("A= %d + B %d = C %d : %d\n", a, b, c, a*b*c)
						return
					}
				}
			}
		}
	}
}
