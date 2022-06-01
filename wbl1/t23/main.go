package main

import "fmt"

func importantOrder(slice []int, s int) []int {
	//Очень медленно
	return append(slice[:s], slice[s+1:]...)
}

func unimportantOrder(s []int, i int) []int {
	//Очень быстро
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(importantOrder(a, 3))
	fmt.Println(unimportantOrder(a, 3))
}
