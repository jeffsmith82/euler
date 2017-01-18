package main

import (
	"fmt"
	"math/big"
)

func main() {

	power := big.NewInt(101)
	nums := big.NewInt(101)
	m := make(map[string]bool)

	for i := big.NewInt(2); i.Cmp(nums) == -1; i.Add(big.NewInt(1), i) {
		for j := big.NewInt(2); j.Cmp(power) == -1; j.Add(big.NewInt(1), j) {
			power := new(big.Int).Exp(i, j, nil)
			m[power.String()] = true
		}
	}
	fmt.Println(m)
	fmt.Println(len(m))
}
