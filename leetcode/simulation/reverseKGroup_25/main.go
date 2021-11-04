package main

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: nil,
		},
	}
	reverseKGroup(node, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	var newHead, end *ListNode
	p := head
	bucket := make([]*ListNode, k)
	for {
		idx := 0
		for i := 0; i < k; i++ {
			if p == nil {
				break
			}
			idx++
			bucket[i] = p
			p = p.Next
		}
		if idx == k {
			for i := k - 1; i >= 0; i-- {
				if newHead == nil {
					newHead = bucket[i]
				}
				if end == nil {
					end = bucket[i]
				} else {
					end.Next = bucket[i]
					end = end.Next
					end.Next = nil
				}
			}
		} else {
			for i := 0; i < idx; i++ {
				if newHead == nil {
					newHead = bucket[i]
				}
				if end == nil {
					end = bucket[i]
				} else {
					end.Next = bucket[i]
					end = end.Next
					end.Next = nil
				}
			}
			return newHead
		}
	}
}
