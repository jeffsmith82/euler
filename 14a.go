package main

import (
	"fmt"
)

func main() {

	max := 0
	term := 0
	counter := 1
	number := 1000000
	//m := make(map[int]int)

	for i := 2; i <= number; i++ {
		counter = 1
		j := i
		for j != 1 {
			//x, ok := m[j]
			//if ok == true {
			//	counter = counter + x - 1
			//	break
			//}
			if j%2 == 0 {
				j = j / 2
			} else {
				j = (j * 3) + 1
			}
			counter++
		}
		//m[i] = counter
		fmt.Println(i, counter)
		if max < counter {
			max = counter
			term = i
		}
	}
	//fmt.Println(m)
	fmt.Println("Term: ", term)
	fmt.Println("max: ", max)
}
