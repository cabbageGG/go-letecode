/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 处理好相加以及进位的问题
	pre := &ListNode{} // 构造一个虚节点
	cur := pre
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum >= 10 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}
		cur.Next = &ListNode{Val: sum} // 多加一个最后的0
		cur = cur.Next
	}
	return pre.Next
}