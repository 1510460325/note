package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	for i := 0; i < n; i++ {
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
	}
	if p2 == nil {
		return p1.Next
	}
	for {
		if p2.Next == nil {
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	p1.Next = p1.Next.Next
	return head
}
