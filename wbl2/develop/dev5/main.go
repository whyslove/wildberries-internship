package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var after, before, context int
	var count, ignoreCase, invert, fixed, lineNum bool
	var outputFile string
	flag.IntVar(&after, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&context, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&count, "c", false, "сжать выход до количества строк")
	flag.BoolVar(&ignoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&invert, "v", false, "выводить НЕ совпадения")
	flag.BoolVar(&fixed, "F", false, "искать точное совпадение со строкой")
	flag.BoolVar(&lineNum, "n", false, "напечатать номер строки")
	flag.StringVar(&outputFile, "output", "", "имя выходного файла")

	flag.Parse()

	inputFilename := flag.Arg(0)
	subStringToFind := flag.Arg(1)

	f, err := os.Open(inputFilename)
	check(err)
	defer f.Close()

	var input []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	result := grep(input, subStringToFind, after, before, context, count, ignoreCase, invert, fixed, lineNum)

	for _, el := range result {
		fmt.Println(el)
	}

	if outputFile != "" {
		f2, err := os.Create(outputFile)
		check(err)
		defer f2.Close()

		w := bufio.NewWriter(f2)
		for _, str := range result {
			w.WriteString(str + "\n")
		}
		w.Flush()

	}
}

func grep(input []string, substring string, after, before, context int, count, ignoreCase, invert, fixed, lineNum bool) []string {

	if before < context {
		before = context
	}
	if after < context {
		after = context
	}

	var result []string
	var accepted = make([]bool, len(input))

	for i, str := range input {
		if checkString(str, substring, ignoreCase, fixed) {
			startInd := i - before
			endInd := i + after
			if startInd < 0 {
				startInd = 0
			}
			if endInd >= len(input) {
				endInd = len(input) - 1
			}
			// fmt.Println(startInd, endInd, i, after, before)
			for j := startInd; j <= endInd; j++ {
				if !accepted[j] {
					accepted[j] = true
				}

			}
		}

	}

	var countLines int

	for i := 0; i < len(input); i++ {
		findMatching := true
		if invert {
			findMatching = false
		}
		if accepted[i] == findMatching {
			if count {
				countLines++
			} else {
				if lineNum {
					result = append(result, (fmt.Sprint(i+1) + " " + input[i]))
				} else {
					result = append(result, input[i])
				}
			}
		}

	}
	if count {
		return []string{fmt.Sprint(countLines)}
	}
	fmt.Println(accepted)
	return result

}

func checkString(str, substr string, ignorecase, fixed bool) bool {
	if ignorecase {
		str = strings.ToLower(str)
	}
	if fixed {
		if str == substr {
			return true
		}
	} else if strings.Contains(str, substr) {
		return true
	}
	return false
}
