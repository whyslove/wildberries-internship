package main

import "fmt"

func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

func main() {
	s1 := "Hello, 世界"
	s2 := "The quick brown 狐 jumped over the lazy 犬"
	s3 := "абырвалг"
	fmt.Println(Reverse(s1))
	fmt.Println(Reverse(s2))
	fmt.Println(Reverse(s3))
}
