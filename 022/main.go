package main

import "fmt"

// 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

// 例如，给出 n = 3，生成结果为：

// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]

var res []string

func generateParenthesis(n int) []string {
	res = []string{}
	backtrace("", n, 0, 0)
	return res
}

func backtrace(S string, n, left, right int) {
	if len(S) >= 2*n {
		res = append(res, S)
	}
	if left < n {
		backtrace(S+"(", n, left+1, right)
	}
	if right < left {
		backtrace(S+")", n, left, right+1)
	}
}

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}
