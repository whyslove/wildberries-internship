package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindAnagrams(input []string) map[string][]string {
	var result = make(map[string][]string)
	var seenWords = make(map[string]bool)
	for _, str := range input {
		str = strings.ToLower(str)
		array := strings.Split(str, "")
		sort.Strings(array)
		if !seenWords[str] {
			result[strings.Join(array, "")] = append(result[strings.Join(array, "")], str)
			seenWords[str] = true
		}

	}
	return result
}

func main() {
	input := []string{"пятка", "пятак", "тяпка", "листок", "столик", "столик", "СЛИТОК"}
	z := FindAnagrams(input)
	for _, strArray := range z {
		if len(strArray) > 1 {
			fmt.Printf("{%s: %s}\n", strArray[0], strArray[1:])
		}
	}
}
