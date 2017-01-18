package main

import (
	"fmt"
	"math/big"
)

func main() {

	//	max := big.NewInt(4000000000000000000000000000000000000000000000000000000)
	fib := big.NewInt(1)
	last := big.NewInt(0)
	temp := big.NewInt(0)
	var counter int64 = 1

	for 1 == 1 {
		if fib == big.NewInt(1) {
			fib.Add(fib, big.NewInt(1))
			continue
		}
		temp.Set(fib)
		fib.Add(last, fib)
		last.Set(temp)
		/*if fib.Cmp(max) == -1 {
			if fib.Mod(fib, big.NewInt(2)) == big.NewInt(0) {
				fmt.Println(fib)
				counter = counter + fib.Int64()
			}
		}*/
		//fmt.Println(counter, fib)
		counter++
		if len(fib.String()) == 1000 {
			fmt.Println(fib)
			fmt.Println(counter)
			return
		}
	}
	fmt.Println("Answer: ", counter)
}
