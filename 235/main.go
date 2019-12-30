package main

// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

// 解法：需要知道，当p，q在一个节点的左右子树时，该节点就是最近公共祖先！！！

//Definition for TreeNode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if q.Val > root.Val && p.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if q.Val < root.Val && p.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	pVal := p.Val
	qVal := q.Val
	for root != nil {
		if qVal > root.Val && pVal > root.Val {
			root = root.Right
		} else if qVal < root.Val && pVal < root.Val {
			root = root.Left
		} else {
			break
		}
	}
	return root
}

func main() {

}
