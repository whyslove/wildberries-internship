package main

import (
	"errors"
	"testing"
)

type TestUnpackingStruct struct {
	input   string
	res     string
	res_err error
}

// var err = errors.New("неправильная строка подана на вход")

func TestUpacking(t *testing.T) {
	var tests = []TestUnpackingStruct{
		{"abcd", "abcd", nil},
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"45", "", errStr},
		{`qwe\4\5`, "qwe45", nil},
		{`qwe\45`, "qwe44444", nil},
		{`qwe\\5`, `qwe\\\\\`, nil},
	}

	for i, test := range tests {
		if res, err := Unpack(test.input); res != test.res || !errors.Is(err, test.res_err) {
			t.Errorf("Error in test %v, inputted %s, expected {%s, %v} get {%s, %v}",
				i, test.input, test.res, test.res_err, res, err)
		}
	}
}
