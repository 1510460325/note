package main

import "fmt"

func main() {
	genList := func(nums ...int) *ListNode {
		head := &ListNode{Val: nums[0]}
		p := head
		for _, i := range nums[1:] {
			p.Next = &ListNode{Val: i}
			p = p.Next
		}
		return head
	}

	printList := func(head *ListNode) {
		for head != nil {
			fmt.Printf("%v ",head.Val)
			head = head.Next
		}
	}
	printList(partition(genList(1,4,3,2,5,2), 3))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	ans := &ListNode{Next: head}
	minEnd := ans
	if minEnd.Next.Val < x {
		minEnd = minEnd.Next
	}
	for head.Next != nil {
		if head.Next.Val >= x {
			head = head.Next
			continue
		}
		if head == minEnd {
			minEnd = minEnd.Next
			head = head.Next
			continue
		}
		tmp := head.Next
		head.Next = head.Next.Next
		tmp.Next = minEnd.Next
		minEnd.Next = tmp
		minEnd = minEnd.Next
	}
	return ans.Next
}
