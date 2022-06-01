package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fields, delim string
	var separated bool
	flag.StringVar(&fields, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&delim, "d", "\t", "разделитель")
	flag.BoolVar(&separated, "s", false, "пропустить не принадлежащие никакому полю")
	flag.Parse()

	if fields == "" {
		fmt.Println("Необходимо выбрать поля")
		os.Exit(1)
	}

	var scanner = bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	result := cut(input, fields, delim, separated)
	for _, v := range result {
		fmt.Println(v)
	}
}

func cut(input []string, fields, delim string, separated bool) []string {
	var result []string
	var fieldsToSearch []int
	//Обработать флаг для полей
	if strings.Contains(fields, "-") {
		for _, el := range strings.Split(fields, "-") {
			integer, err := strconv.Atoi(el)
			if err != nil {
				fmt.Println("Неправильный формат поля")
				os.Exit(1)
			}
			fieldsToSearch = append(fieldsToSearch, integer)
		}
	} else {
		for _, el := range strings.Split(fields, ",") {
			integer, err := strconv.Atoi(el)
			if err != nil {
				fmt.Println("Неправильный формат поля")
				os.Exit(1)
			}
			fieldsToSearch = append(fieldsToSearch, integer)
		}
	}

	for _, str := range input {
		if !strings.Contains(str, delim) {
			if !separated {
				result = append(result, str)
			}
		} else {
			var tempString []string
			_str := strings.Split(str, delim)
			for _, i := range fieldsToSearch {
				if i < len(_str) {
					tempString = append(tempString, _str[i])
				}
			}
			result = append(result, strings.Join(tempString, delim))
		}

	}

	return result
}
