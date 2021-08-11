package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var head, pre *ListNode
	a := 0
	for l1 != nil || l2 != nil || a > 0 {
		sum := a
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		a = sum / 10
		node := &ListNode{Val: sum % 10}
		if head == nil {
			head = node
			pre = node
		} else {
			pre.Next = node
			pre = node
		}
	}

	return head
}
