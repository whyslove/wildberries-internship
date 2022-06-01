package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int) //size is 1
	squares := make(chan int)

	// Var 1
	// a := [5]int{2, 4, 6, 8, 10}
	// go func() {
	// 	for _, val := range a {
	// 		naturals <- val
	// 	}
	// 	close(naturals)
	// }()

	// go func() {
	// 	for val := range naturals {
	// 		squares <- val * val
	// 	}
	// 	close(squares)
	// }()

	// sum := 0
	// for val := range squares {
	// 	sum += val
	// }
	// fmt.Println(sum)

	//Var 2
	a := [5]int{2, 4, 6, 8, 10}
	go func() {
		for _, val := range a {
			naturals <- val
		}
		close(naturals)
	}()

	go func() {
		for val := range naturals {
			squares <- val * val
		}
		close(squares)
	}()

	const N = 4 //number of my cores
	partSums := make(chan int, N)

	for i := 0; i < N; i++ {
		go func() {
			presum := 0
			for value := range squares {
				presum += value
			}
			partSums <- presum
		}()
	}
	var sum int
	for i := 0; i < N; i++ {
		sum += <-partSums
	}
	fmt.Println(sum)
}
