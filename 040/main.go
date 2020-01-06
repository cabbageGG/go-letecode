package main

// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

// candidates 中的每个数字在每个组合中只能使用一次。

// 说明：

// 所有数字（包括目标数）都是正整数。
// 解集不能包含重复的组合。
// 示例 1:

// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 所求解集为:
// [
//   [1, 7],
//   [1, 2, 5],
//   [2, 6],
//   [1, 1, 6]
// ]
// 示例 2:

// 输入: candidates = [2,5,2,1,2], target = 5,
// 所求解集为:
// [
//   [1,2,2],
//   [5]
// ]

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
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
		//注意需要去重
		if i > index && candidates[i-1] == candidates[i] {
			// [1, 1, 2, 5, 6, 7, 10]
			// 0 号索引的 1 ，从后面 6 个元素中搜索
			// 1 号索引也是 1 ，从后面 5 个元素（是上面 6 个元素的真子集）中搜索，这种情况显然上面已经包含
			continue
		}
		nums = append(nums, candidates[i])
		sum = sum + candidates[i]
		// 可重复取值，所以这里传入的下一个index是i
		if !backtrace(candidates, nums, res, i+1, target, sum) {
			break // 这里剪枝
		}
		nums = nums[:len(nums)-1]
		sum = sum - candidates[i]
	}
	return true
}

func main() {
	candidates := []int{1, 1, 2, 5, 6, 7, 10}
	target := 8
	fmt.Println(combinationSum(candidates, target))
}
