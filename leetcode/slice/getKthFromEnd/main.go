package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	p := head
	for i := 0; i < k; i++ {
		p = p.Next
	}
	for p != nil {
		head = head.Next
		p = p.Next
	}
	return head
}

func swapPairs(head *ListNode) *ListNode {
	var pre *ListNode
	p := head
	for p != nil && p.Next != nil {
		if pre != nil {
			pre.Next = p.Next
		} else {
			head = p.Next
		}
		pre = p
		p.Next, p.Next.Next = p.Next.Next, p
		p = p.Next
	}
	return head
}
