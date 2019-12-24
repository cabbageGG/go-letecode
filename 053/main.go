package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum <= 0 {
			sum = nums[i]
		} else {
			if sum+nums[i] < 0 {
				sum = 0
				continue
			} else {
				sum = sum + nums[i]
			}
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {
	a := []int{1, 2, 3, -4, 1}
	fmt.Println(maxSubArray(a))
}
