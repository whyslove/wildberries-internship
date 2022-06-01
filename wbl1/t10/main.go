package main

import (
	"fmt"
)

func main() {
	dict := make(map[int][]float64)

	a := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 19, 21}

	for _, value := range a {
		dict[int(value/10)*10] = append(dict[int(value/10)*10], value)
	}
	fmt.Println(dict)
}
