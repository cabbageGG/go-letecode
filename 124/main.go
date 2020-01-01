package main

// 给定一个非空二叉树，返回其最大路径和。

// 本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
// 示例 1:

// 输入: [1,2,3]

//        1
//       / \
//      2   3

// 输出: 6

//递归解法，保存以每个节点为顶点的路径和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func maxPathSum(root *TreeNode) int {
	res = root.Val
	maxGain(root)
	return res
}

func maxGain(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := maxGain(node.Left)
	if left < 0 {
		left = 0
	}
	right := maxGain(node.Right)
	if right < 0 {
		right = 0
	}
	temp := left + right + node.Val
	if temp > res {
		res = temp
	}
	// 该节点不是顶点
	if left > right {
		return node.Val + left
	}
	return node.Val + right
}

func main() {

}
