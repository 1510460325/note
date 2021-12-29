package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	var pre *Node
	for len(queue) > 0 {
		cur := make([]*Node, 0, len(queue)*2)
		for _, item := range queue {
			if pre != nil {
				pre.Next = item
			}
			pre = item
			if item.Left != nil {
				cur = append(cur, item.Left)
			}
			if item.Right != nil {
				cur = append(cur, item.Right)
			}
		}
		queue, pre = cur, nil
	}
	return root
}
