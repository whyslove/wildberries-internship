package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type NumericSortData struct {
	num int32
	all string
}

func main() {
	var columnSort int
	var numericSort bool
	var reverseOrder bool
	var hideRepeated bool

	flag.IntVar(&columnSort, "k", -1, "колонка для сортировки")
	flag.BoolVar(&numericSort, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&reverseOrder, "r", false, "выводить в обратном порядке")
	flag.BoolVar(&hideRepeated, "u", false, "скрывать неповторяющиеся строки")

	flag.Parse()

	filename1 := flag.Arg(0)
	filename2 := flag.Arg(1)
	f, err := os.Open(filename1)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var inputStrings []string
	for scanner.Scan() {
		inputStrings = append(inputStrings, scanner.Text())
	}

	inputStrings = makeSort(inputStrings, columnSort, numericSort, reverseOrder, hideRepeated)

	for _, str := range inputStrings {
		fmt.Println(str)
	}

	f2, err := os.Create(filename2)
	check(err)
	defer f2.Close()

	w := bufio.NewWriter(f2)
	for _, str := range inputStrings {
		w.WriteString(str + "\n")
	}

	w.Flush()
}

const DELIM = " "

func makeSort(input []string, columnSort int, numericSort, reverseOrder, hideRepeated bool) []string {
	//Порядок выполнения операций в функции
	//Если надо удалить дубли, то удаляем
	//Потом двигаем колонку на первое место (в конце передвинем на правильное место)
	//Потом если numeric то сортируем как numeric
	//Если нет, то дефолт сортировка
	//Если порядок не тот, то переворачиваем
	//Если двигали колонки возвращаем на место

	maxLen := len(input)
	if hideRepeated {
		for i := 0; i < maxLen-1; {
			if input[i] == input[i+1] {
				input[i+1], input[maxLen-1] = input[maxLen-1], input[i+1]
				maxLen--
			}
			i++
		}
		input = input[:maxLen]
	}

	if columnSort != -1 && columnSort > 0 {
		for i := 0; i < len(input); i++ {
			columns := strings.Split(input[i], DELIM)
			if columnSort < len(columns) {
				columns[0], columns[columnSort] = columns[columnSort], columns[0]
			}

			input[i] = strings.Join(columns, " ")
		}

	}

	if numericSort {
		minInt32 := -2147483648
		var inputNum []NumericSortData
		//Create struct for sorting
		for _, el := range input {

			num := 0
			for i := 0; i < len(el); i++ {
				if unicode.IsDigit(rune(el[i])) {
					d, _ := strconv.Atoi(string(el[i]))
					num = num*10 + d
				} else {
					break
				}
			}

			if len(el) == 0 || !unicode.IsDigit(rune(el[0])) {
				num = minInt32
			}

			inputNum = append(inputNum, NumericSortData{num: int32(num), all: el})

		}

		sort.Slice(inputNum, func(i, j int) bool { return inputNum[i].num < inputNum[j].num })
		for i := 0; i < len(inputNum); i++ {
			input[i] = inputNum[i].all
		}
	} else {
		sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })

	}

	if reverseOrder {
		for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
			input[i], input[j] = input[j], input[i]
		}

	}

	if columnSort != -1 && columnSort > 0 {
		for i := 0; i < len(input); i++ {
			columns := strings.Split(input[i], DELIM)
			if columnSort < len(columns) {
				columns[0], columns[columnSort] = columns[columnSort], columns[0]
			}
			input[i] = strings.Join(columns, " ")
		}

	}
	return input
}
