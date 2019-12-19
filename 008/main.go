package main

import "fmt"

func myAtoi(str string) int {
	m := make(map[byte]int, 12)
	num := "0123456789"
	for i := 0; i < len(num); i++ {
		m[num[i]] = i
	}
	sum := 0
	first := false
	sig := 1
	for i := 0; i < len(str); i++ {
		if !(first) && str[i] == ' ' {
			continue
		}
		if !first && str[i] == '-' {
			first = true
			sig = -1
			continue
		}
		if !first && str[i] == '+' {
			first = true
			continue
		}

		if v, ok := m[str[i]]; ok {
			if !first {
				first = true
			}
			sum = sum*10 + v*sig
			if sum < -0x7FFFFFFF-1 {
				sum = -0x7FFFFFFF - 1
				break
			}
			if sum > 0x7FFFFFFF {
				sum = 0x7FFFFFFF
				break
			}
		} else {
			break
		}
	}
	return sum
}

func main() {

	s := "9223372036854775808"
	fmt.Println(myAtoi(s))
}
