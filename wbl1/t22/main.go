package main

import (
	"fmt"
	"math/big"
)

func main() {

	var bignum1, _ = new(big.Int).SetString("218882428714186575617", 0)
	var bignum2, _ = new(big.Int).SetString("218882428714186575617", 0)

	fmt.Println(bignum1.Mul(bignum1, bignum2).Text(10))
}
