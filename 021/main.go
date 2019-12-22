package main

import "fmt"

// 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 示例：

// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
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

func main() {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}
	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}
	l3 := mergeTwoLists(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}
