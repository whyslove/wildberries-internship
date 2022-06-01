package main

import "fmt"

//Donnovan
func main() {
	squares := make(chan int)
	a := [5]int{2, 4, 6, 8, 10}

	go func() {
		for val := range a {
			squares <- val * val
		}
		close(squares)
	}()

	//Concurrently output
	for val := range squares {
		fmt.Println(val)
	}

}
