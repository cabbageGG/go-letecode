package main

import "fmt"

// 给定一个可包含重复数字的序列，返回所有不重复的全排列。

// 示例:

// 输入: [1,1,2]
// 输出:
// [
//   [1,1,2],
//   [1,2,1],
//   [2,1,1]
// ]

// 解法： 此题与上一题的区别在于，需要去重
// 去重的话，先需要知道，什么情况会发生重复
// [1,3,1,2]
// 如上，当1与1交换则相当于重复
// 我们的步骤是：
// 1、选择1，求后面的全排列
// 2、依次让1与后面的数进行交换，交换后，加上剩余的全排列。
// 这样在方案2里面，如果1和1交换，求剩余的全排列。这种情况就与方案1的全排列完全重复。

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	backtrace(0, nums, &res)
	return res
}

func backtrace(first int, nums []int, res *[][]int) {
	if first == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
	}
	m := map[int]bool{}
	for i := first; i < len(nums); i++ {
		if m[nums[i]] {
			continue
		}
		m[nums[i]] = true
		nums[first], nums[i] = nums[i], nums[first]
		backtrace(first+1, nums, res)
		nums[first], nums[i] = nums[i], nums[first]
	}
}

func main() {
	nums := []int{2, 2, 1, 1}
	fmt.Println(permuteUnique(nums))
}
