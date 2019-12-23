package main

import "fmt"

// func removeDuplicates(nums []int) int {
// 	if len(nums) <= 0 {
// 		return 0
// 	}
// 	i := 1
// 	cur := nums[0]
// 	for i < len(nums) {
// 		if cur == nums[i] {
// 			nums = append(nums[:i], nums[i+1:]...)
// 		} else {
// 			cur = nums[i]
// 			i++
// 		}
// 	}
// 	fmt.Println(nums)
// 	return len(nums)
// }
func removeDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	// 双指针算法，让数组后面的元素覆盖掉前面重复的元素
	// 两个指针，分别遍历整个数组，以及保存不重复的元素
	// nums[0,k] 不重复元素数组
	// nums[i, l-1] 带遍历的数组
	i, k := 0, 0
	for i < len(nums) {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
		i++
	}
	fmt.Println(nums[:k+1])
	return k + 1
}
func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
