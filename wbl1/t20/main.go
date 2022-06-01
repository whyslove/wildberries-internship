package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	s1 := "Hello, 世界"
	s2 := "The quick brown 狐 jumped over the lazy 犬"
	s3 := "абырвалг"
	s4 := " "
	s5 := ""
	fmt.Println(reverseWords(s1))
	fmt.Println(reverseWords(s2))
	fmt.Println(reverseWords(s3))
	fmt.Println(reverseWords(s4))
	fmt.Println(reverseWords(s5))
}
