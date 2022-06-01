package main

import "fmt"

func main() {
	//a1 and a2 already sets
	a1 := []int{1, 2, 3, 4, 5, 6}
	a2 := []int{2, 4, 6, 8}
	intersection := make(map[int]int)
	for _, el := range a1 {
		intersection[el] += 1
	}
	for _, el := range a2 {
		intersection[el] += 1
	}
	fmt.Printf("{ ")
	for key, value := range intersection {
		if value > 1 {
			fmt.Printf("%d ", key)
		}

	}
	fmt.Printf("}\n")
}
