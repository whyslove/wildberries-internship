package main

import "fmt"

func main() {
	i := 3
	a := 10
	fmt.Printf("%b\n", a)
	// '|' это побитовое или,
	// a		=10010100101
	// 1<<(I-1) =0000000100
	fmt.Printf("%b\n", a|1<<(i-1))
}
