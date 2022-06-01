package main

import (
	"fmt"
	"unicode"
)

func check(s string) bool {
	dict := make(map[rune]bool)
	for _, symbol := range s {
		symbol = unicode.ToLower(symbol)
		if a := dict[symbol]; !a {
			dict[symbol] = true
		} else {
			return false
		}
	}
	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	fmt.Printf("%v, %v, %v", check(s1), check(s2), check(s3))

}
