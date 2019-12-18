package main

import "fmt"
import _ "github.com/stretchr/testify/assert"

// '''
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// Example:
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].
// '''

// 解题思路
// 使用空间来代替时间。空间来保存已经查证了的元素。

func twoSum(input []int, target int)[]int{
	hasFound := make(map[int]int, len(input))
	for i, v := range input{
		if j, ok := hasFound[target-v]; ok {
			return []int{j,i}
		}
		hasFound[v] = i
	}
	return nil
}

func main(){
	params := []int{1,2,3}
	ans := 5
	fmt.Println(twoSum(params, ans))
}

