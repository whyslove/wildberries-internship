package main

import (
	"strings"
)

//
// var justString string
//BAD
// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

//BETTER
func someFunc() string {
	//1ое не использовать глобальные переременные, которые
	// не используются дальше, потому что они занимают память
	var justString string
	var sb strings.Builder

	//Также возможно (95%), но я не уверен, v ниже будет заставлять
	//хранить в опретивке всю огромную сторку, потому что эта переменная
	//будет ссылаться на часть огромного массива
	//поэтому можем попробовать пересохранить
	v := createHugeString(1 << 10)[:100]
	justString = v[:100]
	sb.WriteString(justString)
	return sb.String()

}

func main() {
	someFunc()
}

func createHugeString(n int) string {
	//неважно что здесь
	return "11111111111111111111"
}
