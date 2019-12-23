package main

import "fmt"

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。

// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

// 你可以假设数组中不存在重复的元素。

// 你的算法时间复杂度必须是 O(log n) 级别。

// 示例 1:

// 输入: nums = [4,5,6,7,0,1,2], target = 0
// 输出: 4
// 示例 2:

// 输入: nums = [4,5,6,7,0,1,2], target = 3
// 输出: -1

func search(nums []int, target int) int {
	// 变种二分查找
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[left] {
			// mid在左边正常序列
			if nums[mid] < target {
				left = mid + 1
			} else {
				if nums[left] > target {
					left = mid + 1
				} else if nums[left] == target {
					return left
				} else {
					right = mid - 1
				}
			}
			// 这里很关键，需要判断mid是在正常序列还是，旋转序列里
			// 此时target有可能在左边或右边
			// 假设在左边，则
		} else {
			// mid在旋转序列
			if nums[mid] > target {
				right = mid - 1
			} else {
				if nums[right] == target {
					return right
				} else if nums[right] < target {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}
	}
	return -1
}

func main() {
	nums := []int{3, 1}
	target := 1
	fmt.Println(search(nums, target))

}
