package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (ret [][]int) {
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		next := make([]*TreeNode, 0, len(queue)*2)
		cur := make([]int, len(queue))
		for i, item := range queue {
			cur[i] = item.Val
			if item.Left != nil {
				next = append(next, item.Left)
			}
			if item.Right != nil {
				next = append(next, item.Right)
			}
		}
		ret = append(ret, cur)
		queue = next
	}
	return ret
}
