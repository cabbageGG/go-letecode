package main

import "fmt"

var res []string
var Map map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	res = []string{}
	if len(digits) > 0 {
		backtrace("", digits)
	}
	return res
}

func backtrace(com, nextDigits string) {
	if len(nextDigits) == 0 {
		res = append(res, com)
		return
	}
	for _, v := range Map[string(nextDigits[0])] {
		backtrace(com+string(v), nextDigits[1:])
	}
}

func main() {
	s := "23"
	fmt.Println(letterCombinations(s))
}
