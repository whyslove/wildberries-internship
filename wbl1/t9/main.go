package main

import "fmt"

//Donnovan
func main() {
	naturals := make(chan int) //size is 1
	squares := make(chan int)
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

	for val := range squares {
		fmt.Println(val)
	}

}
