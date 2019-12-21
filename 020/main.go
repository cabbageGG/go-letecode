package main

import "fmt"

func isValid(s string) bool {
	var left []rune
	m := map[rune]rune{'(': ')', '{': '}', '[': ']'}
	for _, ss := range s {
		if ss == '(' || ss == '{' || ss == '[' {
			left = append(left, ss)
		} else {
			if len(left) <= 0 {
				return false
			}
			if right, ok := m[left[len(left)-1]]; !ok || right != ss {
				return false
			}
			left = left[:len(left)-1]
		}
	}
	if len(left) == 0 {
		return true
	}
	return false
}

func main() {
	s := "{}"
	fmt.Println(isValid(s))
}
