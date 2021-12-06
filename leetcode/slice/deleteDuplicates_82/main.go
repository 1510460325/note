package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	ans := &ListNode{}
	pre := ans
	var consider *ListNode
	for ; head != nil; head = head.Next {
		if consider == nil {
			consider = head
			continue
		}
		// receive
		if consider.Val != head.Val {
			pre.Next = consider
			pre, consider = consider, head
			continue
		}
		// all reject
		for head.Next != nil && head.Next.Val == consider.Val {
			head = head.Next
		}
		consider = nil
		if head == nil {
			break
		}
	}
	if consider != nil {
		pre.Next = consider
		pre = consider
	}
	pre.Next = nil
	return ans.Next
}
