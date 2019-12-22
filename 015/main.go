package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 降维到二元和来处理，由于复杂度在n*2左右，则可直接进行排序
	sort.Ints(nums)
	var rets [][]int
	for i := 0; i < len(nums)-2; i++ {
		// 二元处理 a + b = -v, 范围nms[i+1:]
		// 双指针法
		// 去重？

		// 特别注意去重，一定要在循环的过程中，自动去重不能单独去重。
		// 单独去重，虽然只会增加系数级别的时间复杂度，但是会导致超时
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r]+nums[i] == 0 {
				temp := []int{nums[i], nums[l], nums[r]}
				rets = append(rets, temp)
				// 特别注意去重，一定要在循环的过程中，自动去重不能单独去重。
				// 单独去重，虽然只会增加系数级别的时间复杂度，但是会导致超时
				for l+1 < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r-1 && nums[r-1] == nums[r] {
					r--
				}
				l++
			} else if nums[l]+nums[r]+nums[i] > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return rets
}

func main() {

	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
