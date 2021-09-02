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
