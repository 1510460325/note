package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(reverseBetween(head, 2, 5))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	idx := 0
	ans := &ListNode{}
	p := ans
	for head != nil {
		idx++
		if left <= idx && idx <= right {
			tmp := head.Next
			head.Next = p.Next
			p.Next = head
			head = tmp
			continue
		}
		for p.Next != nil {
			p = p.Next
		}
		p.Next = head
		p = p.Next
		head = head.Next
		p.Next = nil
	}
	return ans.Next
}
