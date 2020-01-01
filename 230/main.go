package main

import "container/list"

// 给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

// 说明：
// 你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	res := inOrder(root)
	//res := inOrder2(root, k)
	if len(res) >= k {
		return res[k-1]
	}
	return -1
}

func inOrder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	arr := []int{}
	left := inOrder(node.Left)
	arr = append(arr, left...)
	arr = append(arr, node.Val)
	right := inOrder(node.Right)
	arr = append(arr, right...)
	return arr
}

//迭代
func inOrder2(node *TreeNode, k int) []int {
	stack := list.New()
	stack.PushBack(node)
	res := []int{}
	for stack.Len() > 0 || node != nil {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}
		v := stack.Back()
		node = v.Value.(*TreeNode)
		stack.Remove(v)
		res = append(res, node.Val)
		if len(res) >= k {
			break
		}
		node = node.Right
	}
	return res
}

func main() {

}
