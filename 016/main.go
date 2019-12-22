package main

import "fmt"

// 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

// 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

// 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

import "sort"

func threeSumClosest(nums []int, target int) int {
	// a + b + c
	sort.Ints(nums)
	if len(nums) < 3 {
		return 0
	}
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 双指针 每次寻找最接近的一个closer的sum
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(res, target) > abs(sum, target) {
				res = sum
			}
			if sum < target {
				l++
			} else if nums[i]+nums[l]+nums[r] > target {
				r--
			} else {
				return target
			}
		}
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}
