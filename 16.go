package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {

	value := int(math.Pow(2, 1000))
	fmt.Println(value)
	stringy := strconv.Itoa(value)
	total := 0
	for i := 0; i < len(stringy); i++ {
		v, err := strconv.Atoi(string(stringy[i]))
		if err != nil {
			log.Fatal(err)
		}
		total = total + v
		fmt.Println(total)
	}
	fmt.Println(total)
}
