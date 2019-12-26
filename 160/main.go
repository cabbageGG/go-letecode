package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	k1, k2 := headA, headB
	flagA, flagB := false, false
	for k1 != k2 {
		if k1.Next == nil {
			if !flagA {
				flagA = true
			} else {
				return nil
			}
			k1 = headB
		} else {
			k1 = k1.Next
		}
		if k2.Next == nil {
			if !flagB {
				flagB = true
			} else {
				return nil
			}
			k2 = headA
		} else {
			k2 = k2.Next
		}
	}
	return k1
}

func main() {

}
