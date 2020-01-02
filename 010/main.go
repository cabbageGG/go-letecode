package main

import "fmt"

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	first_match := len(s) != 0 && (p[0] == s[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (first_match && isMatch(s[1:], p))
	} else {
		return first_match && isMatch(s[1:], p[1:])
	}
}

func main() {
	s := "aaa"
	p := "a*a*"
	fmt.Println(isMatch(s, p))
}
