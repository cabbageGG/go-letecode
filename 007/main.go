package main

import "fmt"

func reverse(x int) int {
	var flag bool
	if x < 0 {
		x = 0 - x
		flag = true
	}
	sum := 0
	for x > 0 {
		sum = sum*10 + x%10
		x = x / 10
	}
	if flag {
		sum = 0 - sum
	}
	if sum < -0x7FFFFFFF || sum > 0x7FFFFFFF-1 {
		sum = 0
	}
	return sum
}

func main() {
	x := -0x7FFFFFFF
	fmt.Println(reverse(x))
}
