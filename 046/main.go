package main

import "fmt"

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	backtrack(0, nums)
	return res
}

func backtrack(first int, nums []int) {
	if first == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		res = append(res, temp)
	}
	for i := first; i < len(nums); i++ {
		nums[first], nums[i] = nums[i], nums[first]
		backtrack(first+1, nums)
		nums[first], nums[i] = nums[i], nums[first]
	}
}

func main() {
	nums := []int{1, 2, 3}
	res1 := permute(nums)
	fmt.Println(res1)
	nums1 := []int{1, 2, 3}
	res2 := permute(nums1)
	fmt.Println(res2)
}
