package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var errStr = errors.New("неправильная строка подана на вход")

func Unpack(input string) (string, error) {
	var sb strings.Builder
	str := []rune(input)

	for i := 0; i < len(str); {

		//Если текущий символ бекслеш применим обработку для бекслешей
		if str[i] == '\\' {
			if i+1 == len(str) {
				return "", errStr
			}
			i += 1

			//Обработка не бекслешей
		} else {
			if unicode.IsDigit(str[i]) {
				return "", errStr
			}
		}

		//Общая часть
		if i+1 == len(str) {
			sb.WriteRune((str[i]))
			i++
		} else if !unicode.IsDigit(str[i+1]) {
			sb.WriteRune(str[i])
			i++
		} else {
			nums := []rune{}
			var j int
			for j = i + 1; j < len(str) && unicode.IsDigit(str[j]); {
				nums = append(nums, str[j])
				j += 1
			}
			cnt, _ := strconv.Atoi(string(nums))
			sb.WriteString(strings.Repeat(string(str[i]), cnt))
			i = j
		}

	}

	return sb.String(), nil

}

func main() {
	str, err := Unpack(`qwe\\5`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}
}
