package main

// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

// candidates 中的数字可以无限制重复被选取。

// 说明：

// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。
// 示例 1:

// 输入: candidates = [2,3,6,7], target = 7,
// 所求解集为:
// [
//   [7],
//   [2,2,3]
// ]
// 示例 2:

// 输入: candidates = [2,3,5], target = 8,
// 所求解集为:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	backtrace(candidates, []int{}, &res, 0, target, 0)
	return res
}

func backtrace(candidates, nums []int, res *[][]int, index, target, sum int) bool {
	if sum == target {
		t := make([]int, len(nums))
		copy(t, nums)
		*res = append(*res, t)
		return true
	} else if sum > target {
		return false
	}
	// sum < target
	for i := index; i < len(candidates); i++ {
		nums = append(nums, candidates[i])
		sum = sum + candidates[i]
		// 可重复取值，所以这里传入的下一个index是i
		if !backtrace(candidates, nums, res, i, target, sum) {
			break // 这里剪枝
		}
		nums = nums[:len(nums)-1]
		sum = sum - candidates[i]
	}
	return true
}

func main() {
	candidates := []int{1, 2, 3, 4}
	target := 4
	fmt.Println(combinationSum(candidates, target))
}
