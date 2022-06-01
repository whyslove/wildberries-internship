package main

import "fmt"

func main() {

	//1
	a := 2
	b := 3
	a, b = b, a
	fmt.Printf("%d, %d \n", a, b)

	//2
	a = 2
	b = 3
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("%d, %d \n", a, b)

	//3
	a = 2
	b = 3
	a = a * b
	b = a / b
	a = a / b
	fmt.Printf("%d, %d \n", a, b)

	//4
	a = 2
	b = 3
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Printf("%d, %d \n", a, b)

	// a = 2
	// b = 3
	// a = (a & b) + (a | b)
	// b = a + (^b) + 1
	// a = a + (^b) + 1
	// fmt.Printf("%d, %d \n", a, b)
}
