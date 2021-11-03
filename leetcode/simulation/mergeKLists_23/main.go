package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	p := head
	for {
		min := -1
		for i := range lists {
			if lists[i] == nil {
				continue
			}
			if min == -1 {
				min = i
			} else if lists[i].Val < lists[min].Val {
				min = i
			}
		}
		if min == -1 {
			break
		}
		p.Next, lists[min] = lists[min], lists[min].Next
		p = p.Next
	}
	return head.Next
}
