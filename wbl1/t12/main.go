package main

import "fmt"

func main() {
	a := make(map[string]bool)
	mas := []string{"cat", "cat", "dog", "cat", "tree"}
	for _, el := range mas {
		a[el] = true
	}
	fmt.Printf("{ ")
	for key := range a {
		fmt.Printf("%s ", key)
	}
	fmt.Printf("}\n")

}
