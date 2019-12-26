package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
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
		return nil
	}
	k2 = head
	for k1 != k2 {
		k1 = k1.Next
		k2 = k2.Next
	}
	return k1
}

func main() {

}
