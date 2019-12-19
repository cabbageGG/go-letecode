package main

import "fmt"

func maxArea(height []int) int {
	// 双指针法
	i, j := 0, len(height)-1
	res := 0
	for i < j {
		var temp int
		if height[i] < height[j] {
			temp = height[i] * (j - i)
			i++
		} else {
			temp = height[j] * (j - i)
			j--
		}
		if temp > res {
			res = temp
		}
	}
	return res
}

func main() {
	v := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(v))
}
