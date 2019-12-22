package main

import "fmt"

// 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

// 示例:

// 输入:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// 输出: 1->1->2->3->4->4->5->6

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// func mergeKLists(lists []*ListNode) *ListNode {
// 	//依次合并两个
// 	var l0 *ListNode
// 	for _, l := range lists {
// 		l0 = mergeTwoLists(l0, l)
// 	}
// 	return l0
// }

// improve
func mergeKLists(lists []*ListNode) *ListNode {
	//归并思路：依次合并两个
	amount := len(lists)
	if amount <= 0 {
		return nil
	}
	interval := 1
	for interval < amount {
		for i := 0; i < amount-interval; i = i + interval*2 {
			lists[i] = mergeTwoLists(lists[i], lists[i+interval])
		}
		interval = interval * 2
	}
	return lists[0]
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 构造一个虚拟头结点，然后依次比较两个链表，把正确的值赋值虚拟节点
	pre := &ListNode{}
	cur := pre
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = &ListNode{Val: l1.Val}
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{Val: l2.Val}
			cur = cur.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return pre.Next
}

// 时间复杂度估算， 每个节点比较（k-1)次，N个节点。O(NK)
// 假设每个链表一样长，则长度为N/k 第二个链表与新链表比较（N/k * 1 + N/k）
// 第三条链表与新链表比较，N/k + （N/k + N/k)        = （N/k * 2 + N/k）
// ...
// 第K条链表与新链表比较                             = （N/k * k-1 + N/k）
// 相加 = N/k * k*(k-1)/2 + N/k = N(k-1)/2 + N/k ~ O(Nk)

// imporve-- O(N*logk)
// 使用归并思想，来两两合并链表

func main() {
	fmt.Println("")
}
