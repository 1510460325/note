package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) (ret [][]int) {
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}
	reverse := false
	for len(queue) > 0 {
		next := make([]*TreeNode, 0, len(queue)*2)
		tmp := make([]int, len(queue))
		start, gap := 0, 1
		if reverse {
			start, gap = len(queue)-1, -1
		}
		for i := range queue {
			cur := queue[i]
			tmp[start] = cur.Val
			start += gap
			if cur.Left != nil {
				next = append(next, cur.Left)
			}
			if cur.Right != nil {
				next = append(next, cur.Right)
			}
		}
		ret = append(ret, tmp)
		reverse = !reverse
		queue = next
	}
	return ret
}
