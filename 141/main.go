package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	k1, k2 := head, head
	flag := false
	for {
		if k2.Next != nil && k2.Next.Next != nil {
			k2 = k2.Next.Next
			k1 = k1.Next
		} else {
			break
		}
		if k1 == k2 {
			flag = true
			break
		}
	}
	if !flag {
		return false
	}
	return true
}

func main() {

}
